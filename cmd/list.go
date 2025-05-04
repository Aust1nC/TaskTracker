package cmd

import (
	"github.com/Aust1nC/TaskTracker/model"
	"github.com/Aust1nC/TaskTracker/service"
	"github.com/spf13/cobra"
)

func ListTasks() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all tasks from the task tracker",
		RunE:  RunListTasksCmd,
	}
}

func RunListTasksCmd(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		status := model.TaskStatusFromString(args[1])
		return service.ListTasks(status)
	}

	return service.ListTasks(model.TaskStatusFromString("all"))
}
