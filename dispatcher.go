package buruh

import "context"

type Dispatcher struct {
	config *Config
	pool   *Pool
	queue  *Queue
	cancel context.CancelFunc
	// signalStop chan bool
}

func New(ctx context.Context, cfg *Config) *Dispatcher {
	q := NewQueue(cfg)
	ctx, cancel := context.WithCancel(ctx)

	p := NewPool(ctx, q, cfg)

	d := &Dispatcher{
		config: cfg,
		pool:   p,
		queue:  q,
		cancel: cancel,
		// signalStop: make(chan bool),
	}

	// d.collect()

	return d
}

func (d *Dispatcher) Dispatch(job Job) {
	d.queue.Enqueue(job)
	// d.pool.Submit(job)
}

func (d *Dispatcher) Debug(t bool) *Dispatcher {
	d.config.Debug = t
	return d
}

// func (d *Dispatcher) collect() {
// 	go func() {
// 		for {
// 			select {
// 			case <-d.signalStop:
// 				return
// 			default:
// 				job, err := d.queue.Dequeue()
// 				if err != nil {
// 					continue
// 				}
// 				d.pool.Submit(job)
// 			}
// 		}
// 	}()
// }

func (d *Dispatcher) Stop() {
	// d.signalStop <- true
	d.cancel()
}
