package cmd

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task-tracker",
		Short: "Task tracker is a CLI tool for managing tasks",
	}

	cmd.AddCommand(NewAddTask())
	cmd.AddCommand(ListTasks())
	return cmd
}
