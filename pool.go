package buruh

import (
	"log"
)

type Pool struct {
	config      *Config
	workers     chan *Worker
	availWorker uint
}

func NewPool(cfg *Config) *Pool {
	p := &Pool{
		config:      cfg,
		workers:     make(chan *Worker, cfg.MaxWorkerNum),
		availWorker: 0,
	}

	return p
}

func (p *Pool) Init() {
	for i := 0; i < int(p.config.MinWorkerNum); i++ {
		if p.config.Debug {
			log.Println("Init new worker")
		}

		w := NewWorker(p.config)
		p.workers <- w
		p.availWorker++
	}
}

func (p *Pool) Get() *Worker {
	for {
		select {
		case w := <-p.workers:
			return w
		default:
			p.addNewWorker()
		}
	}
}

func (p *Pool) addNewWorker() {
	if p.availWorker < p.config.MaxWorkerNum {
		if p.config.Debug {
			log.Println("Add new worker")
		}

		w := NewWorker(p.config)
		p.workers <- w
		p.availWorker++
	}
}
