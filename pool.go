package async

import (
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
