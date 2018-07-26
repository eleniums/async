package async

import (
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
)

func Test_TaskPool_NewTaskPool_Max1_Success(t *testing.T) {
	// act
	pool := NewTaskPool(1)

	// assert
	assert.Equal(t, 1, pool.max)
	assert.NotNil(t, pool.sem)
}

func Test_TaskPool_NewTaskPool_Max2_Success(t *testing.T) {
	// act
	pool := NewTaskPool(2)

	// assert
	assert.Equal(t, 2, pool.max)
	assert.NotNil(t, pool.sem)
}

func Test_TaskPool_NewTaskPool_Max0_Failure(t *testing.T) {
	var pool *TaskPool

	defer func() {
		recover()
		assert.Nil(t, pool)
	}()

	// act
	pool = NewTaskPool(0)

	// assert
	assert.True(t, false)
}

func Test_TaskPool_NewTaskPool_MaxNegative_Failure(t *testing.T) {
	var pool *TaskPool

	defer func() {
		recover()
		assert.Nil(t, pool)
	}()

	// act
	pool = NewTaskPool(-1)

	// assert
	assert.True(t, false)
}

func Test_TaskPool_Run_Success(t *testing.T) {
	// arrange
	startedTask1 := false
	finishedTask1 := false
	task1 := func() error {
		defer func() { finishedTask1 = true }()
		startedTask1 = true
		time.Sleep(time.Millisecond * 100)
		return nil
	}

	startedTask2 := false
	finishedTask2 := false
	task2 := func() error {
		defer func() { finishedTask2 = true }()
		startedTask2 = true
		time.Sleep(time.Millisecond * 200)
		return nil
	}

	startedTask3 := false
	finishedTask3 := false
	task3 := func() error {
		defer func() { finishedTask3 = true }()
		startedTask3 = true
		time.Sleep(time.Millisecond * 200)
		return nil
	}

	pool := NewTaskPool(2)

	// act
	err := pool.Run(task1)
	assert.NoError(t, err)

	err = pool.Run(task2)
	assert.NoError(t, err)

	err = pool.Run(task3)
	assert.NoError(t, err)

	// assert
	assert.True(t, startedTask1)
	assert.True(t, finishedTask1)

	assert.True(t, startedTask2)
	assert.False(t, finishedTask2)

	assert.True(t, startedTask3)
	assert.False(t, finishedTask3)
}

func Test_TaskPool_Wait_Success(t *testing.T) {
	// arrange
	startedTask1 := false
	finishedTask1 := false
	task1 := func() error {
		defer func() { finishedTask1 = true }()
		startedTask1 = true
		time.Sleep(time.Millisecond * 200)
		return nil
	}

	startedTask2 := false
	finishedTask2 := false
	task2 := func() error {
		defer func() { finishedTask2 = true }()
		startedTask2 = true
		time.Sleep(time.Millisecond * 200)
		return nil
	}

	pool := NewTaskPool(2)

	err := pool.Run(task1)
	assert.NoError(t, err)

	err = pool.Run(task2)
	assert.NoError(t, err)

	// act
	err = pool.Wait()

	// assert
	assert.NoError(t, err)

	assert.True(t, startedTask1)
	assert.True(t, finishedTask1)

	assert.True(t, startedTask2)
	assert.True(t, finishedTask2)
}
