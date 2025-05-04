package service

import (
	"encoding/json"
	"fmt"

	"github.com/Aust1nC/TaskTracker/model"
)

func AddTask(description string) error {
	var tasks []model.Task
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var newTaskID int64
	if len(tasks) > 0 {
		lastTask := tasks[len(tasks)-1]
		newTaskID = lastTask.ID + 1
	} else {
		newTaskID = 1
	}

	newTask := model.NewTask(newTaskID, description)

	tasks = append(tasks, *newTask)

	if err := WriteTasksToFile(tasks); err != nil {
		fmt.Println("Error adding a new task:", err)
		return err
	}

	fmt.Printf("\nTask added successfully: %d\n\n", newTaskID)
	return nil
}

func ListTasks(status model.TaskStatus) error {
	var tasks []model.Task
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	filterTasks := filterTasks(tasks, status)

	j, _ := json.MarshalIndent(filterTasks, "", "	")

	fmt.Println(string(j))

	return nil
}

func filterTasks(tasks []model.Task, status model.TaskStatus) []model.Task {
	var filteredTasks []model.Task

	switch status {
	case model.TASK_STATUS_TODO:
		for _, task := range tasks {
			if task.Status == model.TASK_STATUS_TODO {
				filteredTasks = append(filteredTasks, task)
			}
		}
	case model.TASK_STATUS_IN_PROGRESS:
		for _, task := range tasks {
			if task.Status == model.TASK_STATUS_IN_PROGRESS {
				filteredTasks = append(filteredTasks, task)
			}
		}
	case model.TASK_STATUS_COMPLETED:
		for _, task := range tasks {
			if task.Status == model.TASK_STATUS_COMPLETED {
				filteredTasks = append(filteredTasks, task)
			}
		}
	default:
		filteredTasks = tasks
	}
	return filteredTasks
}
