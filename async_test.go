package async

import (
	"context"
	"errors"
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
)

func Test_Run_Success(t *testing.T) {
	// arrange
	task1Completed := false
	task1 := func() error {
		defer func() { task1Completed = true }()
		return nil
	}

	task2Completed := false
	task2 := func() error {
		defer func() { task2Completed = true }()
		return nil
	}

	task3Completed := false
	task3 := func() error {
		defer func() { task3Completed = true }()
		return nil
	}

	// act
	errc := Run(task1, task2, task3)
	err := Wait(errc)

	// assert
	assert.NoError(t, err)
	assert.True(t, task1Completed)
	assert.True(t, task2Completed)
	assert.True(t, task3Completed)
}

func Test_Run_Error(t *testing.T) {
	// arrange
	task1Completed := false
	task1 := func() error {
		defer func() { task1Completed = true }()
		return nil
	}

	task2Completed := false
	task2 := func() error {
		defer func() { task2Completed = true }()
		time.Sleep(time.Millisecond * 200)
		return nil
	}

	task3Completed := false
	task3 := func() error {
		defer func() { task3Completed = true }()
		time.Sleep(time.Millisecond * 100)
		return errors.New("task3 error")
	}

	// act
	errc := Run(task1, task2, task3)
	err := Wait(errc)

	// assert
	assert.Error(t, err)
	assert.True(t, task1Completed)
	assert.False(t, task2Completed)
	assert.True(t, task3Completed)
}

func Test_RunLimited_Success(t *testing.T) {
	// arrange
	count := 0
	task := func() error {
		count++
		return nil
	}

	// act
	err := RunLimited(3, 4, task)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 12, count)
}

func Test_RunLimited_Error(t *testing.T) {
	// arrange
	count := 0
	task := func() error {
		if count < 8 {
			count++
			if count == 8 {
				return errors.New("error")
			}
		}

		return nil
	}

	// act
	err := RunLimited(3, 4, task)

	// assert
	assert.Error(t, err)
	assert.Equal(t, 8, count)
}

func Test_RunForever_Cancel(t *testing.T) {
	// arrange
	ctx, cancel := context.WithCancel(context.Background())

	count := 0
	task := func() error {
		if count >= 10 {
			cancel()
		}

		count++

		return nil
	}

	// act
	errc := RunForever(ctx, 2, task)
	err := Wait(errc)

	// assert
	assert.Error(t, err)
	assert.True(t, count >= 10)
}

func Test_RunForever_Timeout(t *testing.T) {
	// arrange
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	task := func() error {
		return nil
	}

	// act
	errc := RunForever(ctx, 2, task)
	err := Wait(errc)

	// assert
	assert.Error(t, err)
}

func Test_RunForever_Deadline(t *testing.T) {
	// arrange
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond*100))
	defer cancel()

	task := func() error {
		return nil
	}

	// act
	errc := RunForever(ctx, 2, task)
	err := Wait(errc)

	// assert
	assert.Error(t, err)
}

func Test_RunForever_Error(t *testing.T) {
	// arrange
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	count := 0
	task := func() error {
		if count >= 10 {
			return errors.New("task error")
		}

		count++

		return nil
	}

	// act
	errc := RunForever(ctx, 2, task)
	err := Wait(errc)

	// assert
	assert.Error(t, err)
	assert.True(t, count >= 10)
}
