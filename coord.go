package main

import (
    "sync"
)

type Coordinator struct {
    scheduler Scheduler
    workers   []Worker
    wg        sync.WaitGroup
}

func NewCoordinator(numWorkers int) *Coordinator {
    coordinator := &Coordinator{
        scheduler: NewScheduler(),
        workers:   make([]Worker, numWorkers),
    }

    for i := range coordinator.workers {
        coordinator.workers[i] = NewWorker(i+1, coordinator.scheduler.TasksChannel(), &coordinator.wg)
    }

    return coordinator
}


func (c *Coordinator) Start() {
    c.scheduler.Start()
    for i := range c.workers {
        c.workers[i].Start()
    }
}

func (c *Coordinator) AddTask(task Task) {
    c.scheduler.AddTask(task)
}

func (c *Coordinator) Stop() {
    c.scheduler.Stop()
    for i := range c.workers {
        c.workers[i].Stop()
    }
}

func (c *Coordinator) Wait() {
    c.wg.Wait()
}
