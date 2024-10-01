package main

import (
	"fmt"
	"os"
	"strconv"
	"task-tracker/helpers"
	"task-tracker/repository"
	"task-tracker/services"
	"task-tracker/utils"
)

func main() {
	arguments := os.Args[1:]
	fmt.Println(arguments)
	taskService := services.NewTaskService(repository.NewTaskRepository())
	if utils.IsValidArguments(arguments) {
		option := arguments[0]
		utils.GETOPTIONSELECTED(option)
		switch option {
		case utils.ADDTASK:
			task := helpers.CreateTask(arguments)
			taskService.AddTask(task)
		case utils.UPDATETASK:
			if helpers.ValidCommand(arguments, utils.ACCEPTCANTARGUMENTSUPDATE) {
				newDescription := arguments[2]
				idTask, _ := strconv.Atoi(arguments[1])
				taskService.UpdateTask(utils.DESCRIPTIONFIELD, newDescription, idTask)
			} else {
				fmt.Println(utils.TOLESSARGUMENTS)
			}
		case utils.DELETETASK:
			if helpers.ValidCommand(arguments, utils.ACCEPTCANTARGUMENTSDELETE) {
				idTask, _ := strconv.Atoi(arguments[1])
				taskService.DeleteTask(float64(idTask))
			} else {
				fmt.Println(utils.TOLESSARGUMENTS)
			}
		case utils.MARKPROGRESSTASK:
			if helpers.ValidCommand(arguments, utils.ACCEPTCANTARGUMENTSMARKTASK) {
				idTask, _ := strconv.Atoi(arguments[1])
				statusProgress := strconv.Itoa(utils.MARKPROGRESS)
				taskService.UpdateTask(utils.STATUSFIELD, statusProgress, idTask)
			} else {
				fmt.Println(utils.TOLESSARGUMENTS)
			}
		case utils.MARKDONETASK:
		default:
			fmt.Println(utils.OPTIONNOTVALID)
		}
	} else {
		fmt.Println(utils.DONTHAVEOPTION)
	}
	os.Exit(0)
}
