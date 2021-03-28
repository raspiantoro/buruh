package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/raspiantoro/buruh"
)

var fn = func(id int, wg *sync.WaitGroup) buruh.Task {
	return func(ctx context.Context) {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(1)
		time.Sleep(time.Duration(n) * time.Second)

		wKey := ctx.Value(buruh.CtxWorkerKey).(string)
		jKey := ctx.Value(buruh.CtxJobKey).(string)

		fmt.Printf("Method #%d is executing by job: %s with worker: %s\n", id, jKey, wKey)
		wg.Done()
	}
}

func main() {
	ctx := context.Background()
	dispatcher := buruh.New(ctx, &buruh.Config{
		QueueSize:    5,
		MaxWorkerNum: 3,
		MinWorkerNum: 3,
	})

	numOfJob := 20
	wg := sync.WaitGroup{}
	wg.Add(numOfJob)

	for i := 1; i <= numOfJob; i++ {
		job := buruh.NewJob(fn(i, &wg), true)
		dispatcher.Dispatch(job)
	}

	wg.Wait()

	dispatcher.Stop()
}
