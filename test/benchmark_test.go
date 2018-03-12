package test

import (
	"testing"

	"github.com/eleniums/async"
)

func Benchmark_Run(b *testing.B) {
	task := func() error {
		return nil
	}

	for i := 0; i < b.N; i++ {
		async.Run(task, task, task)
	}
}
