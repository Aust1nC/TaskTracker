package model

import "time"

type TaskStatus string

const (
	TASK_STATUS_TODO        TaskStatus = "todo"
	TASK_STATUS_IN_PROGRESS TaskStatus = "in-progress"
	TASK_STATUS_COMPLETED   TaskStatus = "completed"
)

type Task struct {
	ID          int64      `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func TaskStatusFromString(status string) TaskStatus {
	switch status {
	case "todo":
		return TASK_STATUS_TODO
	case "in-progress":
		return TASK_STATUS_IN_PROGRESS
	case "completed":
		return TASK_STATUS_COMPLETED
	default:
		return "all"
	}
}

func NewTask(id int64, description string) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Status:      TASK_STATUS_TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
