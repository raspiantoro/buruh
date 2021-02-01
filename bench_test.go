package buruh_test

import (
	"context"
	"testing"

	"github.com/raspiantoro/buruh"
)

var testFn = func(a int, b int) buruh.Task {
	return func(ctx context.Context) {
		c := a + b

		ctx = context.WithValue(ctx, "calc-result", c)
	}
}

func BenchmarkWithoutBuruh(b *testing.B) {
	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		// perform the operation we're analyzing
		fn := testFn(i, i+1)
		// job := buruh.NewJob(fn)
		// go job.Do(ctx)
		go fn(ctx)
	}
}

func BenchmarkWithBuruh(b *testing.B) {
	dispatcher := buruh.New(&buruh.Config{
		MaxWorkerNum: 100000,
		MinWorkerNum: 5000,
	})

	for i := 0; i < b.N; i++ {
		// perform the operation we're analyzing
		fn := testFn(i, i+1)
		job := buruh.NewJob(fn, false)
		dispatcher.Dispatch(job)
	}
}
