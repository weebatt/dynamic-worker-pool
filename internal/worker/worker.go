package worker

import "fmt"

type Job struct {
	ID      int
	Name    string
	Payload any
}

type Worker struct {
	ID       int
	JobQueue chan Job
	QuitChan chan bool
}

func NewWorker(id int) Worker {
	return Worker{
		ID:       id,
		JobQueue: make(chan Job),
		QuitChan: make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			select {
			case job := <-w.JobQueue:
				// Process the job
				fmt.Printf("Worker %d: Processing job %d\n", w.ID, job.ID)
			case <-w.QuitChan:
				// Receive stop signal
				fmt.Printf("Worker %d stopping\n", w.ID)
				return
			}
		}
	}()
}
