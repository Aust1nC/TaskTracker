package cmd

import (
	"fmt"
	"strconv"

	"github.com/Aust1nC/TaskTracker/service"
	"github.com/spf13/cobra"
)

func DeleteTask() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a task from the tracker",
		RunE:  RunDeleteTaskCmd,
	}
}

func RunDeleteTaskCmd(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a task id")
	}

	taskIDInt, err := strconv.ParseInt(args[0], 10, 32)
	if err != nil {
		return err
	}

	return service.DeleteTask(taskIDInt)
}
