package helpers

import (
	"fmt"
	"task-tracker/models"
	"task-tracker/utils"
)

func ShowAllTaks(tasks []models.Task) {
	if utils.HaveTask(tasks) {
		for i, task := range tasks {
			fmt.Printf("Task N° %v\n", i+1)
			task.ShowTask()
		}
	} else {
		fmt.Printf(utils.DONTHAVETASKS)
	}
}
