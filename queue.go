package buruh

import (
	"errors"
	"sync/atomic"
)

type item struct {
	job     *Job
	enqueue bool
}

type CircularQueue struct {
	jobs              []item
	readPos, writePos int64
	count             int64
}

func NewCircularQueue(size int) *CircularQueue {
	return &CircularQueue{
		jobs:     make([]item, size),
		writePos: 0,
	}
}

func (c *CircularQueue) Dequeue() (job *Job, err error) {

	for {
		if !c.jobs[c.readPos].enqueue {
			err = errors.New("queue is empty")
			return
		}

		var newValue int64
		oldValue := c.readPos

		if oldValue == int64(len(c.jobs)-1) {
			newValue = 0
		} else {
			newValue = oldValue + 1
		}

		job = c.jobs[oldValue].job

		if job == nil {
			continue
		}

		swapped := atomic.CompareAndSwapInt64(&c.readPos, oldValue, newValue)
		if swapped {
			c.jobs[oldValue].enqueue = false

			return
		}

	}

}

func (c *CircularQueue) Enqueue(job *Job) (err error) {

	for {
		if c.jobs[c.writePos].enqueue {
			err = errors.New("queue is full")
			return
		}

		var newValue int64
		oldValue := c.writePos

		if oldValue == int64(len(c.jobs)-1) {
			newValue = 0
		} else {
			newValue = oldValue + 1
		}

		swapped := atomic.CompareAndSwapInt64(&c.writePos, oldValue, newValue)
		if swapped {
			c.jobs[oldValue].job = job
			c.jobs[oldValue].enqueue = true

			atomic.AddInt64(&c.count, 1)

			return
		}

	}

}
