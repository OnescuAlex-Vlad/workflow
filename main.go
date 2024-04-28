package main

import (
	"fmt"
)

// Task represents a single task in a workflow
type Task struct {
	Name        string
	HTTPRequest string
}

// Workflow represents a workflow with a sequence of tasks
type Workflow struct {
	Name  string
	Tasks []Task
}


func main() {
	var workflow Workflow = Workflow{Name: "Test WF",
		Tasks: []Task{{
			Name: "Task 1",
			HTTPRequest: "http://localhost:8080",
		}}}

	fmt.Println(workflow)
}
