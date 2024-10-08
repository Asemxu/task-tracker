package services

import (
	"task-tracker/models"
	"task-tracker/repository"
	"task-tracker/utils"
)

type TaskServiceFunc interface {
	AddTask()
	UpdateTask()
	DeleteTask()
	MarkTask()
	ListTasks()
}

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return TaskService{repo: repo}
}

func (task TaskService) AddTask(taskValue models.Task) {
	taskValue.Id = float64(len(task.ListTasks(utils.NOTQUERY, utils.EMPTYVALUE, utils.EMPTYVALUE)) + 1)
	task.repo.AddDB(taskValue)
}

func (task TaskService) UpdateTask(fieldName string, value string, idTask int) {
	task.repo.UpdateDB(fieldName, value, idTask)
}

func (task TaskService) DeleteTask(idTask float64) {
	task.repo.DeleteDB(idTask)
}

func (task TaskService) MarkTask(isMarkProgress bool, isMarkDone bool) {

}

func (task TaskService) ListTasks(isQuery bool, field string, valueQuery string) []models.Task {
	return task.repo.GetDB(isQuery, field, valueQuery)
}
