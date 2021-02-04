package buruh

import (
	"context"
	"log"
)

type Pool struct {
	config      *Config
	workers     chan *Worker
	jobsQueue   chan chan Job
	availWorker uint
}

func NewPool(cfg *Config) *Pool {

	p := &Pool{
		config:      cfg,
		workers:     make(chan *Worker, cfg.MaxWorkerNum),
		jobsQueue:   make(chan chan Job, 100),
		availWorker: 0,
	}

	return p
}

func (p *Pool) Init(ctx context.Context) {
	for i := 0; i < int(p.config.MinWorkerNum); i++ {
		if p.config.Debug {
			log.Println("Init new worker")
		}

		p.addNewWorker(ctx)
	}
}

func (p *Pool) submit(job chan Job) {
	p.jobsQueue <- job
}

func (p *Pool) addNewWorker(ctx context.Context) {
	if p.availWorker < p.config.MaxWorkerNum {
		if p.config.Debug {
			log.Println("Add new worker")
		}

		w := NewWorker(p.config)
		p.workers <- w
		//go w.Start(ctx, p.jobsQueue)
		p.availWorker++
	}
}
