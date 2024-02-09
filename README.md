# Task Scheduler Application in Go

This is a task scheduler application written in Go, aimed at scheduling tasks and executing them concurrently using multiple workers.

## Overview

The project consists of several files:

- `main.go`: Contains the main entry point for the application.
- `coordinator.go`: Defines the Coordinator struct responsible for managing the scheduling of tasks and workers.
- `scheduler.go`: Defines the Scheduler struct responsible for queuing tasks.
- `worker.go`: Defines the Worker struct responsible for executing tasks.

## Getting Started

### Prerequisites

- Go programming language installed on your machine.

### Installation

1. Clone the repository:

    ```bash
    git clone <repository_url>
    ```

2. Navigate to the project directory:

    ```bash
    cd task-scheduler-go
    ```

3. Build the application:

    ```bash
    go build
    ```

4. Run the executable:

    ```bash
    ./task-scheduler-go
    ```

## Usage

The application schedules tasks and executes them concurrently using multiple workers. To use the task scheduler, follow these steps:

1. Define tasks in the `main.go` file by creating instances of the `Task` struct and adding them to the coordinator.
2. Start the coordinator to begin scheduling and executing tasks.
3. Wait for all tasks to be completed.
