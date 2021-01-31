package buruh

import "time"

type Config struct {
	MaxWorkerNum  uint
	MinWorkerNum  uint
	MaxWorkerLife time.Duration
	Debug         bool
}
