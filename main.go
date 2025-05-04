package main

import (
	"fmt"

	"github.com/Aust1nC/TaskTracker/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
}
