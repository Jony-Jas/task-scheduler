package main

import (
    "fmt"
)

type Task struct {
    ID          int
    Description string
}

func main() {
    coordinator := NewCoordinator(3) // creating a coordinator with 3 workers
    coordinator.Start()
    defer coordinator.Stop()

    tasks := []Task{
        {ID: 1, Description: "Task 1"},
        {ID: 2, Description: "Task 2"},
        {ID: 3, Description: "Task 3"},
        {ID: 4, Description: "Task 4"},
        {ID: 5, Description: "Task 5"},
    }

    for _, task := range tasks {
        coordinator.AddTask(task)
    }

    coordinator.Wait()
    fmt.Println("All tasks completed.")
}
