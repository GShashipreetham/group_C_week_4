package main

import (
    "fmt"
)

var tasks []string

func addTask(task string) {
    tasks = append(tasks, task)
}

func listTasks() {
    fmt.Println("To-Do List:")
    for i, task := range tasks {
        fmt.Printf("%d: %s\n", i+1, task)
    }
}