package async

import (
	"context"

	"golang.org/x/sync/semaphore"
)

type TaskPool struct {
	max int
	sem *semaphore.Weighted
}

func NewTaskPool(max int) *TaskPool {
	return &TaskPool{
		max: max,
		sem: semaphore.NewWeighted(int64(max)),
	}
}

func (p *TaskPool) Process(task Task) error {
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

func (p *TaskPool) Wait() error {
	for i := 0; i < p.max; i++ {
		err := p.sem.Acquire(context.Background(), 1)
		if err != nil {
			return err
		}
	}
	for i := 0; i < p.max; i++ {
		p.sem.Release(1)
	}

	return nil
}
