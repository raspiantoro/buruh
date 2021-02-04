package buruh_test

import (
	"context"
	"testing"
	"time"

	"github.com/raspiantoro/buruh"
)

var testFn = func(a int, b int, done chan<- bool) buruh.Task {
	return func(ctx context.Context) {
		ctx = context.WithValue(ctx, "calc-result", a+b)
		done <- true
	}
}

func BenchmarkWithoutBuruh(b *testing.B) {
	ctx := context.Background()

	done := make(chan bool)

	for i := 0; i < b.N; i++ {
		// perform the operation we're analyzing
		// fn := testFn(i, i+1, done)
		// job := buruh.NewJob(fn)
		// go job.Do(ctx)
		// go fn(ctx)
		go testFn(i, i+1, done)(ctx)
		<-done
	}
}

var globalCtx = context.Background()
var dispatcher = buruh.New(globalCtx, &buruh.Config{
	MaxWorkerNum: 500000,
	MinWorkerNum: 500000,
	CoolingTime:  200 * time.Millisecond,
	WarmTime:     500 * time.Millisecond,
	BackoffTime:  200 * time.Millisecond,
})

func BenchmarkWithBuruh(b *testing.B) {

	done := make(chan bool)

	for i := 0; i < b.N; i++ {
		// perform the operation we're analyzing
		// fn := testFn(i, i+1, done)
		// job := buruh.NewJob(fn, false)
		job := buruh.NewJob(testFn(i, i+1, done), false)
		dispatcher.Dispatch(job)
		// dispatcher.Dispatch(buruh.NewJob(testFn(i, i+1, done), false))
		<-done
	}
}
