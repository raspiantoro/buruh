package buruh

import (
	"context"
	"log"
	"time"

	"github.com/rs/xid"
)

type CtxKey string

var (
	CtxWorkerKey CtxKey = "worker-key"
	CtxJobKey    CtxKey = "job-key"
)

type Worker struct {
	config    *Config
	ID        xid.ID
	startTime time.Time
}

func NewWorker(cfg *Config) *Worker {
	uid := xid.New()

	if cfg.Debug {
		log.Printf("Spawn new worker, id: %s", uid.String())
	}

	return &Worker{
		ID:        uid,
		config:    cfg,
		startTime: time.Now(),
	}
}

func (w *Worker) Start(ctx context.Context, queue *CircularQueue) {
	heartBeat := time.NewTicker(w.config.HearbeatRate)
	for range heartBeat.C {
		select {
		case <-ctx.Done():
			return
		default:
			job, err := queue.Dequeue()
			if err != nil {
				if w.config.Debug {
					log.Println(err)
				}

				continue
			}

			if job == nil {
				continue
			}

			if w.config.Debug {
				log.Printf("Execute job: %s, with worker: %s", job.ID.String(), w.ID.String())
			}

			ctx := context.Background()
			ctx = context.WithValue(ctx, CtxWorkerKey, w.ID.String())
			ctx = context.WithValue(ctx, CtxJobKey, job.ID.String())

			job.Do(ctx)
			time.Sleep(w.config.CoolingTime)

			if w.config.Debug {
				log.Printf("Finish job: %s, with worker: %s", job.ID.String(), w.ID.String())
			}
		}

	}
}
