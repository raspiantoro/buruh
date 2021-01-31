package buruh

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

type Task func(ctx context.Context)

type Job struct {
	ID   uuid.UUID
	task Task
}

func NewJob(task Task) (job *Job) {
	uid := uuid.NewV4()

	job = &Job{
		ID:   uid,
		task: task,
	}

	return
}

func (j *Job) Do(ctx context.Context) {
	j.task(ctx)
}
