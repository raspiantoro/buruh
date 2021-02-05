package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/raspiantoro/buruh"
)

var fn = func(id int, wg *sync.WaitGroup) buruh.Task {
	return func(ctx context.Context) {
		wKey := ctx.Value(buruh.CtxWorkerKey).(string)
		jKey := ctx.Value(buruh.CtxJobKey).(string)

		fmt.Sprintf("Method #%d is executing by job: %s with worker: %s\n", id, jKey, wKey)

		// time.Sleep(5 * time.Millisecond)
		wg.Done()
	}
}

func main() {
	var (
		start  time.Time
		elapse time.Duration
	)

	numOfJob := 10000000
	wg := sync.WaitGroup{}

	ctx := context.Background()

	wg.Add(numOfJob)
	cfg := &buruh.Config{
		MaxWorkerNum: 10,
		MinWorkerNum: 10,
		CoolingTime:  1 * time.Nanosecond,
		HearbeatRate: 1 * time.Nanosecond,
		// WarmTime:     500 * time.Microsecond,
		// BackoffTime:  100 * time.Microsecond,
	}

	loopStart := time.Now()
	dispatcher := buruh.New(ctx, cfg)
	loopElapse := time.Since(loopStart)

	fmt.Println("Time consume for pool creation: ", loopElapse)

	start = time.Now()

	for i := 1; i <= numOfJob; i++ {
		dispatcher.Dispatch(buruh.NewJob(fn(i, &wg), false))
	}

	wg.Wait()

	elapse = time.Since(start)
	fmt.Println("elapse time with buruh: ", elapse)

	time.Sleep(1 * time.Second)

	wg.Add(numOfJob)
	start = time.Now()

	for i := 1; i <= numOfJob; i++ {
		go func(i int) {
			ctx = context.WithValue(ctx, buruh.CtxWorkerKey, string(i))
			ctx = context.WithValue(ctx, buruh.CtxJobKey, string(i))

			f := fn(i, &wg)
			f(ctx)
		}(i)
	}

	wg.Wait()

	elapse = time.Since(start)
	fmt.Println("elapse time without buruh: ", elapse)

}
