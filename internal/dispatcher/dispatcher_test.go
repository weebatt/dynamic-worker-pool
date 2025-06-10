package dispatcher

import (
	"sync"
	"testing"
	"time"
)

func TestDispatcher_AddRemoveWorkers(t *testing.T) {
	d := NewDispatcher(1)
	if len(d.workers) != 1 {
		t.Fatalf("expected 1 worker, got %d", len(d.workers))
	}

	d.AddWorker()
	if len(d.workers) != 2 {
		t.Fatalf("expected 2 workers, got %d", len(d.workers))
	}

	d.RemoveWorker()
	if len(d.workers) != 1 {
		t.Fatalf("expected 1 worker after removal, got %d", len(d.workers))
	}
}

func TestDispatcher_JobProcessing(t *testing.T) {
	var wg sync.WaitGroup
	jobCount := 5
	wg.Add(jobCount)

	d := NewDispatcher(0)
	d.ProcessFunc = func(id int, job string) {
		wg.Done()
	}
	d.AddWorker()
	d.AddWorker()

	for i := 0; i < jobCount; i++ {
		d.Dispatch("test job")
	}
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("timeout waiting for jobs to be processed")
	}
}

func TestDispatcher_Race(t *testing.T) {
	d := NewDispatcher(2)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if i%2 == 0 {
				d.AddWorker()
			} else {
				d.RemoveWorker()
			}
			d.Dispatch("race job")
		}(i)
	}
	wg.Wait()
}
