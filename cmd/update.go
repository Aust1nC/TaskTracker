package cmd

import (
	"fmt"
	"strconv"

	"github.com/Aust1nC/TaskTracker/service"
	"github.com/spf13/cobra"
)

func UpdateTaskDescription() *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "Update a task from the task list",
		RunE:  RunUpdateTaskDescriptionTaskCmd,
	}
}

func RunUpdateTaskDescriptionTaskCmd(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("please provide a task id and the new description")
	}

	taskID := args[0]
	taskIDInt, err := strconv.ParseInt(taskID, 10, 32)
	if err != nil {
		return err
	}

	newDescription := args[1]
	return service.UpdateTaskDescription(taskIDInt, newDescription)
}

func UpdateTaskStatusToInProgress() *cobra.Command {
	return &cobra.Command{
		Use:   "mark-in-progress",
		Short: "Update the status of a task to in-progress",
		RunE:  RunUpdateTaskStatusToInProgress,
	}
}

func RunUpdateTaskStatusToInProgress(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a task id")
	}

	taskID := args[0]
	taskIDInt, err := strconv.ParseInt(taskID, 10, 32)
	if err != nil {
		return err
	}

	return service.UpdateTaskStatusToInProgress(taskIDInt)
}

func UpdateTaskStatusToCompleted() *cobra.Command {
	return &cobra.Command{
		Use:   "mark-done",
		Short: "Update the status of a task to completed",
		RunE:  RunUpdateTaskStatusToCompleted,
	}
}

func RunUpdateTaskStatusToCompleted(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a task id")
	}

	taskID := args[0]
	taskIDInt, err := strconv.ParseInt(taskID, 10, 32)
	if err != nil {
		return err
	}

	return service.UpdateTaskStatusToCompleted(taskIDInt)
}
