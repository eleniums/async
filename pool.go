package async

import (
	"context"

	"golang.org/x/sync/semaphore"
)

// TaskPool limits the number of concurrent tasks being processed to a given max.
type TaskPool struct {
	max int
	sem *semaphore.Weighted
}

// NewTaskPool creates a new task pool that will limit concurrent tasks to max.
func NewTaskPool(max int) *TaskPool {
	return &TaskPool{
		max: max,
		sem: semaphore.NewWeighted(int64(max)),
	}
}

// Run will block until there is available capacity and then execute the given task.
func (p *TaskPool) Run(task Task) error {
	err := p.sem.Acquire(context.Background(), 1)
	if err != nil {
		return err
	}

	go func() {
		defer p.sem.Release(1)
		task()
	}()

	return nil
}

// Wait until all tasks have finished processing.
func (p *TaskPool) Wait() error {
	// acquire all available slots in semaphore
	for i := 0; i < p.max; i++ {
		err := p.sem.Acquire(context.Background(), 1)
		if err != nil {
			return err
		}
	}

	// all tasks have completed; release the semaphore
	for i := 0; i < p.max; i++ {
		p.sem.Release(1)
	}

	return nil
}
