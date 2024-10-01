package utils

import "task-tracker/models"

func HaveTask(tasks []models.Task) bool {
	return len(tasks) > EMPTYTASKS
}
