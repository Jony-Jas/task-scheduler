package main

import (
    "sync"
)

type Scheduler struct {
    tasks      []Task
    tasksMutex sync.Mutex
    tasksCh    chan Task
}

func NewScheduler() Scheduler {
    return Scheduler{
        tasksCh: make(chan Task),
    }
}

func (s *Scheduler) Start() {
    go func() {
        for {
            select {
            case task := <-s.tasksCh:
                s.tasksMutex.Lock()
                s.tasks = append(s.tasks, task)
                s.tasksMutex.Unlock()
            }
        }
    }()
}

func (s *Scheduler) Stop() {
    close(s.tasksCh)
}

func (s *Scheduler) AddTask(task Task) {
    s.tasksCh <- task
}

func (s *Scheduler) TasksChannel() <-chan Task {
    return s.tasksCh
}
