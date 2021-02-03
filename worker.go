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
	jobQueue  chan Job
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
		jobQueue:  make(chan Job, 100),
		startTime: time.Now(),
	}
}

func (w *Worker) Start(ctx context.Context, jobCh chan chan Job) {
	for {
		select {
		case jch := <-<-jobCh:
			w.jobQueue <- jch
		case job := <-w.jobQueue:
			// if w.config.Debug {
			// 	log.Printf("Execute job: %s, with worker: %s", job.ID.String(), w.ID.String())
			// }

			// ctx := context.Background()
			// ctx = context.WithValue(ctx, CtxWorkerKey, w.ID.String())
			// ctx = context.WithValue(ctx, CtxJobKey, job.ID.String())

			job.Do(ctx)

			// if w.config.Debug {
			// 	log.Printf("Finish job: %s, with worker: %s", job.ID.String(), w.ID.String())
			// }
		case <-ctx.Done():
			return
		default:
			continue
		}

	}

}