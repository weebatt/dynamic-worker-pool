package dispatcher

import (
	"dynamic-worker-pool/internal/worker"
	"sync"
)

type Dispatcher struct {
	JobQueue    chan string
	workers     []*worker.Worker
	workerCount int
	mu          sync.Mutex
	ProcessFunc func(id int, job string)
}

func NewDispatcher(initialWorkers int) *Dispatcher {
	d := &Dispatcher{
		JobQueue:    make(chan string, 100),
		workers:     make([]*worker.Worker, 0),
		workerCount: 0,
	}
	for i := 0; i < initialWorkers; i++ {
		d.AddWorker()
	}
	return d
}

func (d *Dispatcher) AddWorker() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.workerCount++
	w := worker.NewWorker(d.workerCount)
	w.ProcessFunc = d.ProcessFunc
	d.workers = append(d.workers, w)
	w.StartWithSharedJobQueue(d.JobQueue)
}

func (d *Dispatcher) RemoveWorker() {
	d.mu.Lock()
	defer d.mu.Unlock()
	if len(d.workers) == 0 {
		return
	}
	w := d.workers[len(d.workers)-1]
	w.Stop()
	d.workers = d.workers[:len(d.workers)-1]
}

func (d *Dispatcher) Dispatch(job string) {
	d.JobQueue <- job
}
