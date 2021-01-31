package buruh

import (
	"log"
)

type Dispatcher struct {
	config     *Config
	jobs       *Queue
	pool       *Pool
	stopSignal chan bool
}

func New(cfg *Config) *Dispatcher {
	q := NewQueue(cfg)
	p := NewPool(cfg)

	d := &Dispatcher{
		config:     cfg,
		jobs:       q,
		pool:       p,
		stopSignal: make(chan bool),
	}

	p.Init()
	d.run()

	return d
}

func (d *Dispatcher) Dispatch(job *Job) {
	d.jobs.Enqueue(job)
}

func (d *Dispatcher) Debug(t bool) *Dispatcher {
	d.config.Debug = t
	return d
}

func (d *Dispatcher) run() {
	go func() {
		for {
			select {
			case <-d.stopSignal:
				return
			default:
				err := d.exec()
				if err != nil {
					continue
				}

			}
		}
	}()

}

func (d *Dispatcher) exec() (err error) {
	job, err := d.jobs.Dequeue()
	if err != nil {
		return
	}

	if d.config.Debug {
		log.Println("Waiting for available worker")
	}

	worker := d.pool.Get()

	go worker.Start(job, d.pool.workers)

	return
}

func (d *Dispatcher) Stop() {
	d.stopSignal <- true
}
