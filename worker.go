package buruh

import (
	"context"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

type CtxKey string

var (
	CtxWorkerKey CtxKey = "worker-key"
	CtxJobKey    CtxKey = "job-key"
)

type Worker struct {
	config    *Config
	ID        uuid.UUID
	startTime time.Time
}

func NewWorker(cfg *Config) *Worker {
	uid := uuid.NewV4()

	if cfg.Debug {
		log.Printf("Spawn new worker, id: %s", uid.String())
	}

	return &Worker{
		ID:        uid,
		config:    cfg,
		startTime: time.Now(),
	}
}

func (w *Worker) Start(job *Job, ch chan<- *Worker) {
	if w.config.Debug {
		log.Printf("Execute job: %s, with worker: %s", job.ID.String(), w.ID.String())
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, CtxWorkerKey, w.ID.String())
	ctx = context.WithValue(ctx, CtxJobKey, job.ID.String())

	job.Do(ctx)

	if w.config.Debug {
		log.Printf("Finish job: %s, with worker: %s", job.ID.String(), w.ID.String())
	}

	ch <- w
}
