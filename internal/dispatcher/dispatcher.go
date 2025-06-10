package dispatcher

import (
	"dynamic-worker-pool/internal/worker"
)

type Dispatcher struct {
	WorkerPool chan chan worker.Job
	maxWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan worker.Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, maxWorkers: maxWorkers}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		worker := worker.NewWorker(i + 1)
		worker.Start()
	}
}

func (d *Dispatcher) Dispatch(job worker.Job) {
	go func() {
		// Assign job to an available worker's queue
		workerChannel := <-d.WorkerPool
		workerChannel <- job
	}()
}
