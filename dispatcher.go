package buruh

import (
	"context"
	"time"
)

type Dispatcher struct {
	config *Config
	pool   *Pool
	queue  *CircularQueue
	cancel context.CancelFunc
}

func New(ctx context.Context, cfg *Config) *Dispatcher {
	q := NewCircularQueue(int(cfg.QueueSize))
	ctx, cancel := context.WithCancel(ctx)

	if cfg.HearbeatRate == 0 {
		cfg.HearbeatRate = time.Nanosecond * 5
	}

	p := NewPool(ctx, q, cfg)

	d := &Dispatcher{
		config: cfg,
		pool:   p,
		queue:  q,
		cancel: cancel,
	}

	return d
}

func (d *Dispatcher) Dispatch(job *Job) {
	for {
		err := d.queue.Enqueue(job)
		if err == nil {
			return
		}
	}

}

func (d *Dispatcher) Debug(t bool) *Dispatcher {
	d.config.Debug = t
	return d
}

func (d *Dispatcher) Stop() {
	d.cancel()
}
