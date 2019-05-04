package task

import "time"

type Worker struct {
	ticker   *time.Ticker
	executor *Executor
}

func NewWorker(interval time.Duration, executor *Executor) *Worker {
	return &Worker{time.NewTicker(interval * time.Second), executor}
}

func (worker *Worker) StartWorker() {
	for {
		select {
		case <-worker.ticker.C:
			go worker.executor.Start()
		}
	}
}

func Start() {
	executor := NewExecutor(3, true, VideoClearDispatch, VideoClearExecutor)

	worker := NewWorker(3, executor)

	go worker.StartWorker()
}
