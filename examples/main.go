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
		n := rand.Intn(2)
		time.Sleep(time.Duration(n) * time.Second)

		wKey := ctx.Value(buruh.CtxWorkerKey).(string)
		jKey := ctx.Value(buruh.CtxJobKey).(string)

		fmt.Printf("Method #%d is executing by job: %s with worker: %s\n", id, jKey, wKey)
		wg.Done()
	}
}

func main() {
	dispatcher := buruh.New(&buruh.Config{
		MaxWorkerNum: 10,
		MinWorkerNum: 5,
	})

	numOfJob := 100
	wg := sync.WaitGroup{}
	wg.Add(numOfJob)

	for i := 1; i <= numOfJob; i++ {
		job := buruh.NewJob(fn(i, &wg))
		dispatcher.Dispatch(job)
	}

	wg.Wait()

	dispatcher.Stop()
}
