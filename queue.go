package buruh

// Queue defince struct for job queueing
type Queue struct {
	config *Config
	jobs   []*Job
}

// NewQueue job queueing constructor
func NewQueue(cfg *Config) *Queue {
	return &Queue{
		config: cfg,
		jobs:   []*Job{},
	}
}

// Enqueue add new job to queue
func (q *Queue) Enqueue(job *Job) {
	q.jobs = append(q.jobs, job)
}

// Dequeue retrieve first job from queue
func (q *Queue) Dequeue() (job *Job, err error) {
	if len(q.jobs) == 0 {
		err = errEmptyQueue
		return
	}
	job = q.jobs[0:1][0]
	q.jobs = q.jobs[1:]
	return
}

// First get first job element in queue
func (q *Queue) First() (job *Job, err error) {
	if len(q.jobs) == 0 {
		err = errEmptyQueue
		return
	}
	job = q.jobs[0]
	return
}

// Last get last job element in queue
func (q *Queue) Last() (job *Job, err error) {
	if len(q.jobs) == 0 {
		err = errEmptyQueue
		return
	}
	job = q.jobs[len(q.jobs)-1]
	return
}

// Show get all job element in queue
func (q *Queue) Show() []*Job {
	return q.jobs
}
