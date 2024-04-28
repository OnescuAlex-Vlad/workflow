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

// executeTask simulates the execution of a task by printing a message
func executeTask(task Task) {
	fmt.Printf("Executing task: %s\n", task.Name)
}

// executeWorkflow executes a workflow by sequentially executing each task
func executeWorkflow(workflow Workflow) {
	fmt.Printf("Executing workflow: %s\n", workflow.Name)
	for _, task := range workflow.Tasks {
		executeTask(task)
	}
}

func main() {
	var workflow Workflow = Workflow{Name: "Test WF",
		Tasks: []Task{{
			Name: "Task 1",
			HTTPRequest: "http://localhost:8080",
		}}}

	executeWorkflow(workflow)
}