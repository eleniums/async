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
	if max <= 0 {
		panic("max must be a value of >= 1")
	}

	return &TaskPool{
		max: max,
		sem: semaphore.NewWeighted(int64(max)),
	}
}

// Run will block until there is available capacity and then execute the given task. Cancelling the context will stop the task from being started.
func (p *TaskPool) Run(ctx context.Context, task Task) <-chan error {
	errc := make(chan error, 1)

	err := p.sem.Acquire(ctx, 1)
	if err != nil {
		errc <- err
		close(errc)
		return errc
	}

	go func() {
		defer p.sem.Release(1)
		defer close(errc)

		err = task()
		if err != nil {
			errc <- err
		}
	}()

	return errc
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
	p.sem.Release(int64(p.max))

	return nil
}
