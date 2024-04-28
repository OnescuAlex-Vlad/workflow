package main

import (
	"fmt"
	"sync"
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

	var workflows = map[string]Workflow{
		"Workflow-1": {
			Name: "Workflow-1",
			Tasks: []Task{
				{Name: "Task-1", HTTPRequest: "Fetch document"},
				{Name: "Task-2", HTTPRequest: "Extract content"},
				{Name: "Task-3", HTTPRequest: "Post-process content"},
			},
		},
		"Workflow-2": {
			Name: "Workflow-2",
			Tasks: []Task{
				{Name: "Task-1", HTTPRequest: "Aggregate data"},
				{Name: "Task-2", HTTPRequest: "Transform aggregated data"},
			},
		},
	}

	// Execute the workflows concurrently
	var wg sync.WaitGroup
	for _, wf := range workflows {
		wg.Add(1)

		go func(w Workflow) {
			defer wg.Done()
			executeWorkflow(w)
		}(wf)
	}

	wg.Wait()
}
