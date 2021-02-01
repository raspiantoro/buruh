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
		// wKey := ctx.Value(buruh.CtxWorkerKey).(string)
		// jKey := ctx.Value(buruh.CtxJobKey).(string)

		// fmt.Printf("Method #%d is executing by job: %s with worker: %s\n", id, jKey, wKey)

		wg.Done()
	}
}

func main() {
	var (
		start  time.Time
		elapse time.Duration
	)

	ctx := context.Background()

	numOfJob := 500000
	wg := sync.WaitGroup{}
	wg.Add(numOfJob)

	dispatcher := buruh.New(ctx, &buruh.Config{
		MaxWorkerNum: 10000,
		MinWorkerNum: 10000,
		// Debug:        true,
	})
	defer dispatcher.Stop()
	start = time.Now()

	for i := 1; i <= numOfJob; i++ {
		job := buruh.NewJob(fn(i, &wg), false)
		dispatcher.Dispatch(job)
	}

	wg.Wait()

	elapse = time.Since(start)
	fmt.Println(elapse)

	wg.Add(numOfJob)
	start = time.Now()

	for i := 1; i <= numOfJob; i++ {
		go func(i int) {
			f := fn(i, &wg)
			f(ctx)
		}(i)
	}

	wg.Wait()

	elapse = time.Since(start)
	fmt.Println(elapse)

}
