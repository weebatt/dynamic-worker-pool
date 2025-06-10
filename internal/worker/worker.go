package worker

import (
	"fmt"
)

type Worker struct {
	ID          int
	JobChan     chan string
	quit        chan struct{}
	ProcessFunc func(id int, job string)
}

func NewWorker(id int) *Worker {
	return &Worker{
		ID:      id,
		JobChan: make(chan string),
		quit:    make(chan struct{}),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			select {
			case job := <-w.JobChan:
				if w.ProcessFunc != nil {
					w.ProcessFunc(w.ID, job)
				} else {
					fmt.Printf("Worker %d processing: %s\n", w.ID, job)
				}
			case <-w.quit:
				fmt.Printf("Worker %d stopping\n", w.ID)
				return
			}
		}
	}()
}

// New method: worker reads directly from shared job queue
func (w *Worker) StartWithSharedJobQueue(sharedJobQueue <-chan string) {
	go func() {
		for {
			select {
			case job := <-sharedJobQueue:
				if w.ProcessFunc != nil {
					w.ProcessFunc(w.ID, job)
				} else {
					fmt.Printf("Worker %d processing: %s\n", w.ID, job)
				}
			case <-w.quit:
				fmt.Printf("Worker %d stopping\n", w.ID)
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	close(w.quit)
}
