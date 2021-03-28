package buruh

import "fmt"

var (
	errQueueEmpty = fmt.Errorf("Queue is empty")
	errJobTaken   = fmt.Errorf("Job already taken")
	errNilTask    = fmt.Errorf("Task is nil")
	errNilQueue   = fmt.Errorf("Queue is nil")
	errQueueFull  = fmt.Errorf("Queue is full")
)

var (
	errListEmpty = fmt.Errorf("List is empty")
	errListFull  = fmt.Errorf("List is full")
)
