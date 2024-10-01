package helpers

import (
	"fmt"
	"task-tracker/models"
)

func ShowAllTaks(tasks []models.Task) {
	for i, task := range tasks {
		fmt.Printf("Task N° %v\n", i+1)
		task.ShowTask()
	}
}
