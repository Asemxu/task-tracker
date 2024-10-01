package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"task-tracker/helpers"
	"task-tracker/models"
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
			//ADD TASKS
			task := helpers.CreateTask(arguments)
			taskService.AddTask(task)
		case utils.UPDATETASK:
			if helpers.ValidCommand(arguments, utils.ACCEPTCANTARGUMENTSUPDATE) {
				//UPDATE TASK
				newDescription := arguments[2]
				idTask, _ := strconv.Atoi(arguments[1])
				taskService.UpdateTask(utils.DESCRIPTIONFIELD, newDescription, idTask)
			} else {
				fmt.Println(utils.TOLESSARGUMENTS)
			}
		case utils.DELETETASK:
			if helpers.ValidCommand(arguments, utils.ACCEPTCANTARGUMENTSDELETE) {
				// DELETE TASK
				idTask, _ := strconv.Atoi(arguments[1])
				taskService.DeleteTask(float64(idTask))
			} else {
				fmt.Println(utils.TOLESSARGUMENTS)
			}
		case utils.MARKPROGRESSTASK:
			if helpers.ValidCommand(arguments, utils.ACCEPTCANTARGUMENTSMARKTASK) {
				//MARK IN PROGRESS TASKS
				idTask, _ := strconv.Atoi(arguments[1])
				statusProgress := strconv.Itoa(utils.MARKPROGRESS)
				taskService.UpdateTask(utils.STATUSFIELD, statusProgress, idTask)
			} else {
				fmt.Println(utils.TOLESSARGUMENTS)
			}
		case utils.MARKDONETASK:
			if helpers.ValidCommand(arguments, utils.ACCEPTCANTARGUMENTSMARKTASK) {
				//MARK DONE TASK
				idTask, _ := strconv.Atoi(arguments[1])
				statusProgress := strconv.Itoa(utils.MARKDONE)
				taskService.UpdateTask(utils.STATUSFIELD, statusProgress, idTask)
			} else {
				fmt.Println(utils.TOLESSARGUMENTS)
			}
		case utils.LISTTASKS:
			if helpers.ValidCommand(arguments, utils.ACCEPTCANTARGUMENTSLIST) {
				//GET ALL TASKS
				tasks := taskService.ListTasks(utils.NOTQUERY, utils.EMPTYVALUE, utils.EMPTYVALUE)
				helpers.ShowAllTaks(tasks)
			} else if helpers.ValidCommand(arguments, utils.ACCEPTCANTARGUMENTSLISTQUERY) {
				//GET BY STATUS
				optionList := arguments[1]
				var taskShow []models.Task
				switch optionList {
				case utils.LISTTASKSDONE:
					taskShow = taskService.ListTasks(utils.ISQUERY, strings.ToLower(utils.STATUSFIELD), utils.LISTDONE)
				case utils.LISTTASKSTODO:
					taskShow = taskService.ListTasks(utils.ISQUERY, strings.ToLower(utils.STATUSFIELD), utils.LISTTODO)
				case utils.LISTTASKSINPROGRESS:
					taskShow = taskService.ListTasks(utils.ISQUERY, strings.ToLower(utils.STATUSFIELD), utils.LISTPROGRESS)
				default:
					fmt.Println(utils.OPTIONNOTVALID)
				}
				helpers.ShowAllTaks(taskShow)

			} else {
				fmt.Println(utils.TOLESSARGUMENTS)
			}
		default:
			fmt.Println(utils.OPTIONNOTVALID)
		}
	} else {
		fmt.Println(utils.DONTHAVEOPTION)
	}
	os.Exit(0)
}
