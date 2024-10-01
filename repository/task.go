package repository

import (
	"fmt"
	"reflect"
	"strings"
	"task-tracker/models"
	"task-tracker/utils"
	"time"

	"github.com/davecgh/go-spew/spew"
	db "github.com/sonyarouje/simdb"
)

type TaskRepositoryFunc interface {
	AddDB()
	GetSingleDB()
	GetAllDB()
	UpdateDB()
	DeleteDB()
	GetDB()
}

func NewTaskRepository() TaskRepository {
	driver, err := db.New("data")
	if err != nil {
		panic(err)
	}
	return TaskRepository{db: *driver}
}

type TaskRepository struct {
	db db.Driver
}

func (repo TaskRepository) AddDB(task models.Task) {
	fmt.Println(utils.ADDINGTAKS)
	err := repo.db.Insert(task)
	if err != nil {
		panic(err)
	}
}

func (repo TaskRepository) GetSingleDB(idTask int) models.Task {
	var taskFounded models.Task
	err := repo.db.Open(models.Task{}).Where("id", "=", idTask).First().AsEntity(&taskFounded)
	if err != nil {
		panic(err)
	}
	return taskFounded
}

func (repo TaskRepository) GetAllDB(isQuery bool, field string, valueQuery string) []models.Task {
	var tasks []models.Task
	var err error
	if isQuery {
		err = repo.db.Open(models.Task{}).Where(field, "=", valueQuery).Get().AsEntity(&tasks)
	} else {
		err = repo.db.Open(models.Task{}).Get().AsEntity(&tasks)
	}
	if err != nil {
		if strings.Contains(err.Error(), utils.DONTFOUNDTASK) {
			return tasks
		}
		panic(err)
	}
	return tasks
}

func (repo TaskRepository) UpdateDB(fieldName string, value string, idTask int) {
	fmt.Print(utils.UPDATINGTASK)
	fmt.Println(idTask)
	newValue := reflect.ValueOf(value).String()
	var taskFounded = repo.GetSingleDB(idTask)
	tempValueTask := reflect.ValueOf(&taskFounded).Elem()
	tempValue := reflect.ValueOf(&newValue).Elem()
	fieldValue := tempValueTask.FieldByName(fieldName)
	fielValueUpdated := tempValueTask.FieldByName(utils.UPDATEDATFIELD)

	fielValueUpdated.SetString(time.Now().String())
	fieldValue.Set(tempValue)
	spew.Dump(taskFounded)
	err := repo.db.Update(taskFounded)
	if err != nil {
		panic(err)
	}
	fmt.Println(utils.SUCCESSUPDATE)
}

func (repo TaskRepository) DeleteDB(idTask float64) {
	taskDelete := models.Task{
		Id: idTask,
	}
	err := repo.db.Delete(taskDelete)

	if err != nil {
		fmt.Printf(utils.CANTNOTDELETETASK)
		fmt.Println(int(idTask))
		panic(err)
	}

	fmt.Print(utils.SUCCESSDELETE)
	fmt.Print(idTask)
}

func (repo TaskRepository) GetDB(isQuery bool, field string, valueQuery string) []models.Task {
	fmt.Println(utils.GETTINGALLTASKS)
	var tasks = repo.GetAllDB(isQuery, field, valueQuery)
	return tasks
}
