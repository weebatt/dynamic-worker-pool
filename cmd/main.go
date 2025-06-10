package main

import (
	"dynamic-worker-pool/internal/dispatcher"
	"dynamic-worker-pool/internal/worker"
	"time"
)

func main() {
	maxWorkers := 3
	dispatcher := dispatcher.NewDispatcher(maxWorkers)
	dispatcher.Run()

	for i := 1; i <= 5; i++ {
		job := worker.Job{ID: i, Payload: "Job payload"}
		dispatcher.Dispatch(job)
	}

	time.Sleep(2 * time.Second)
}
