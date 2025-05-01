package model

import "time"

type TaskStatus string

const (
	TASK_STATUS_TODO        TaskStatus = "todo"
	TASK_STATUS_IN_PROGRESS TaskStatus = "in-progress"
	TASK_STATUS_COMPLETED   TaskStatus = "completed"
)

type Task struct {
	ID          string     `json:"id"`
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
