package main

import (
	"dynamic-worker-pool/internal/dispatcher"
	"fmt"
	"time"
)

func main() {
	d := dispatcher.NewDispatcher(2)

	// Send initial jobs
	for i := 1; i <= 5; i++ {
		d.Dispatch(fmt.Sprintf("Job %d payload", i))
	}

	time.Sleep(1 * time.Second)

	// Dynamically add a worker
	fmt.Println("Adding a worker...")
	d.AddWorker()

	// Send more jobs
	for i := 6; i <= 8; i++ {
		d.Dispatch(fmt.Sprintf("Job %d payload", i))
	}

	time.Sleep(1 * time.Second)

	// Dynamically remove a worker
	fmt.Println("Removing a worker...")
	d.RemoveWorker()

	// Send more jobs
	for i := 9; i <= 10; i++ {
		d.Dispatch(fmt.Sprintf("Job %d payload", i))
	}

	time.Sleep(2 * time.Second)
}
