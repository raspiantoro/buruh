package buruh

import (
	"time"
)

type Config struct {
	MaxWorkerNum  uint
	MinWorkerNum  uint
	MaxWorkerLife time.Duration
	CoolingTime   time.Duration
	WarmTime      time.Duration
	BackoffTime   time.Duration
	Debug         bool
}
