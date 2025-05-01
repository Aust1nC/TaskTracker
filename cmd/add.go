package cmd

import (
	"errors"

	"github.com/Aust1nC/TaskTracker/service"
	"github.com/spf13/cobra"
)

func NewAddTask() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add a task to the task list",
		RunE:  RunAddNewTaskCmd,
	}
}

func RunAddNewTaskCmd(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("task description is required")
	}
	description := args[0]
	return service.AddTask(description)
}
