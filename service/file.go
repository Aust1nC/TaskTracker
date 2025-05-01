package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/Aust1nC/TaskTracker/model"
)

func getFilePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
	}
	return path.Join(cwd, "tasks.json")
}

func ReadTasksFromFile() ([]model.Task, error) {
	filePath := getFilePath()

	// Check if file exists, if not, create a new one.
	_, err := os.Stat(filePath)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist, creating a new one...")
		file, err := os.Create(filePath)

		os.WriteFile(filePath, []byte("[]"), os.ModeAppend.Perm())

		if err != nil {
			fmt.Println("Failed to create the file")
			return nil, err
		}

		defer file.Close()
		return []model.Task{}, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}

	var tasks []model.Task

	if err = json.NewDecoder(file).Decode(&tasks); err != nil {
		fmt.Println("Error decoding file:", err)
		return nil, err
	}

	defer file.Close()

	return tasks, nil
}
