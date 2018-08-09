package async

import (
	"errors"
	"sync"
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

func Test_RunForever_Success(t *testing.T) {
	// arrange
	var wg sync.WaitGroup
	wg.Add(12)
	count := 0
	task := func() error {
		if count < 12 {
			defer wg.Done()
			count++
		}

		return nil
	}

	// act
	go RunForever(1, task)
	wg.Wait()

	// assert
	assert.Equal(t, 12, count)
}

func Test_RunForever_Error(t *testing.T) {
	// arrange
	count := 0
	task := func() error {
		if count < 12 {
			count++
			if count == 12 {
				return errors.New("error")
			}
		}

		return nil
	}

	// act
	err := RunForever(1, task)

	// assert
	assert.Error(t, err)
	assert.Equal(t, 12, count)
}
