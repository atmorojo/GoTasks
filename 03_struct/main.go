package main

// Masih agak OOP minded. Semoga ke depannya bisa berkurang

import (
	"fmt"
	"time"
)

type Task struct {
	Todo      string
	CreatedAt time.Time
}

func (t *Task) NewTask(newTask string) {
	t.Todo = newTask
	t.CreatedAt = time.Now()
}

func (t Task) TaskCreated() string {
	return t.CreatedAt.Format(time.RFC822)
}

func main() {
	var task_01 Task
	task_01.NewTask("Kerjakan tugas Go")

	fmt.Printf("Task %q dibuat pada %q\n", task_01.Todo, task_01.TaskCreated())
}
