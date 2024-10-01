package helpers

import (
	"strconv"
	"task-tracker/models"
	"task-tracker/utils"
	"time"
)

func CreateTask(arguments []string) models.Task {

	var description = utils.EMPTYVALUE

	if len(arguments) == 1 {
		description = utils.DESCRIPTIONDEFAULT
	} else {
		description = arguments[1]
	}
	return models.Task{
		//where use random id to more secutiry
		// Id:          utils.GetId(),
		Status:      strconv.Itoa(utils.MARKCREATED),
		CreatedAt:   time.Now().String(),
		UpdatedAt:   "",
		Description: description,
	}
}
