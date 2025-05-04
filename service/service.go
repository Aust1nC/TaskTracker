package service

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/Aust1nC/TaskTracker/model"
)

func AddTask(description string) error {
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
	j, _ := json.MarshalIndent(newTask, "", " ")
	fmt.Println(string(j))

	return nil
}

func UpdateTaskDescription(id int64, description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("Task not found, creating a new task...")
		AddTask(description)
		return nil
	}

	var updatedTask []model.Task
	var existTask model.Task
	for _, task := range tasks {
		if task.ID == id {
			task.Description = description
			task.UpdatedAt = time.Now()
			existTask = task
		}
		updatedTask = append(updatedTask, task)
	}

	if err := WriteTasksToFile(updatedTask); err != nil {
		return err
	}

	j, _ := json.MarshalIndent(existTask, "", " ")
	fmt.Println(string(j))

	return nil
}

func UpdateTaskStatus(id int64, status model.TaskStatus) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var updatedTask []model.Task
	var existTask model.Task
	for _, task := range tasks {
		if task.ID == id {
			task.Status = status
			task.UpdatedAt = time.Now()
			existTask = task
		}
		updatedTask = append(updatedTask, task)
	}

	if err := WriteTasksToFile(updatedTask); err != nil {
		return err
	}

	j, _ := json.MarshalIndent(existTask, "", " ")
	fmt.Println(string(j))

	return nil
}

func UpdateTaskStatusToInProgress(id int64) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var updatedTask []model.Task
	var existTask model.Task
	for _, task := range tasks {
		if task.ID == id {
			task.Status = model.TASK_STATUS_IN_PROGRESS
			task.UpdatedAt = time.Now()
			existTask = task
		}
		updatedTask = append(updatedTask, task)
	}

	if err := WriteTasksToFile(updatedTask); err != nil {
		return err
	}

	j, _ := json.MarshalIndent(existTask, "", " ")
	fmt.Println(string(j))

	return nil
}

func UpdateTaskStatusToCompleted(id int64) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var updatedTask []model.Task
	var existTask model.Task
	for _, task := range tasks {
		if task.ID == id {
			task.Status = model.TASK_STATUS_COMPLETED
			task.UpdatedAt = time.Now()
			existTask = task
		}
		updatedTask = append(updatedTask, task)
	}

	if err := WriteTasksToFile(updatedTask); err != nil {
		return err
	}

	j, _ := json.MarshalIndent(existTask, "", " ")
	fmt.Println(string(j))

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

	// Sort by UpdatedAt descending (newest first)
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].UpdatedAt.After(tasks[j].UpdatedAt)
	})

	filteredTasks := filterTasks(tasks, status)

	j, _ := json.MarshalIndent(filteredTasks, "", "	")

	fmt.Println(string(j))

	return nil
}

func DeleteTask(id int64) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var updatedTask []model.Task
	for _, task := range tasks {
		if task.ID != id {
			updatedTask = append(updatedTask, task)
		}
	}

	if len(tasks) == len(updatedTask) {
		fmt.Println("Task to delete not found:", err)
		return err
	}

	if err := WriteTasksToFile(updatedTask); err != nil {
		return err
	}

	j, _ := json.MarshalIndent(updatedTask, "", " ")
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
