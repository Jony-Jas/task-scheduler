package main

import (
    "fmt"
    "sync"
    "time"
)

type Worker struct {
    id     int
    taskCh <-chan Task
    wg     *sync.WaitGroup
}

func NewWorker(id int, taskCh <-chan Task, wg *sync.WaitGroup) Worker {
    return Worker{
        id:     id,
        taskCh: taskCh,
        wg:     wg,
    }
}

func (w Worker) Start() {
    go func() {
        for task := range w.taskCh {
            fmt.Printf("Worker %d is performing task: %s\n", w.id, task.Description)
            time.Sleep(time.Second) // Simulating task execution time
            w.wg.Done()
        }
    }()
}

func (w Worker) Stop() {
    // Nothing to stop for this simple example
}
