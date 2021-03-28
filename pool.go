package buruh

import (
	"context"
	"log"
)

type Pool struct {
	config      *Config
	jobsQueue   *CircularQueue
	availWorker uint
}

func NewPool(ctx context.Context, jobsQueue *CircularQueue, cfg *Config) *Pool {

	p := &Pool{
		config:      cfg,
		jobsQueue:   jobsQueue,
		availWorker: 0,
	}

	p.init(ctx)

	return p
}

func (p *Pool) init(ctx context.Context) {
	for i := 0; i < int(p.config.MinWorkerNum); i++ {
		if p.config.Debug {
			log.Println("Init new worker")
		}

		p.addNewWorker(ctx)
	}

}

func (p *Pool) addNewWorker(ctx context.Context) {
	if p.availWorker < p.config.MaxWorkerNum {
		if p.config.Debug {
			log.Println("Add new worker")
		}

		w := NewWorker(p.config)
		go w.Start(ctx, p.jobsQueue)
		p.availWorker++
	}
}
