package buruh

import (
	"time"
)

type Config struct {
	QueueSize     uint
	MaxWorkerNum  uint
	MinWorkerNum  uint
	MaxWorkerLife time.Duration
	CoolingTime   time.Duration
	HearbeatRate  time.Duration
	BackoffTime   time.Duration
	Debug         bool
}
