package buruh

import "context"

type Dispatcher struct {
	config *Config
	// jobs   *Queue
	jobs       chan Job
	pool       *Pool
	stopSignal chan bool
	cancel     context.CancelFunc
}

func New(cfg *Config) *Dispatcher {
	// q := NewQueue(cfg)
	p := NewPool(cfg)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	d := &Dispatcher{
		config: cfg,
		// jobs:   q,
		jobs:       make(chan Job, 100),
		pool:       p,
		stopSignal: make(chan bool),
		cancel:     cancel,
	}

	p.Init(ctx)
	d.run()

	return d
}

func (d *Dispatcher) Dispatch(job Job) {
	// d.jobs.Enqueue(job)
	d.jobs <- job
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
				d.cancel()
				return
			default:
				d.pool.jobsQueue <- d.jobs
			}
		}
	}()

}

func (d *Dispatcher) Stop() {
	d.stopSignal <- true
}
