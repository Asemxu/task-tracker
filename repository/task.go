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
	fmt.Println("Agregando una tarea...")
	err := repo.db.Insert(task)
	if err != nil {
		fmt.Println("1")
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

func (repo TaskRepository) GetAllDB() []models.Task {
	var tasks []models.Task
	err := repo.db.Open(models.Task{}).Get().AsEntity(&tasks)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return tasks
		}
		panic(err)
	}
	return tasks
}

func (repo TaskRepository) UpdateDB(fieldName string, value string, idTask int) {
	fmt.Print("Actualizando la tarea con ID: ")
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
	fmt.Println("Actualizado con éxito")
}

func (repo TaskRepository) DeleteDB(idTask float64) {
	taskDelete := models.Task{
		Id: idTask,
	}
	err := repo.db.Delete(taskDelete)

	if err != nil {
		fmt.Printf("No se logro eliminar el Tas con el id ")
		fmt.Println(int(idTask))
		panic(err)
	}

	fmt.Print("Se ah borrado el task con el Id ")
	fmt.Print(idTask)
}

func (repo TaskRepository) GetDB() []models.Task {
	fmt.Println("Obteniendo todos los tasks...")
	var tasks = repo.GetAllDB()
	return tasks
}

// func (repo TaskRepository) UpdateMark(idTask int, statusTask int) {
// 	fmt.Print("Actualizando la tarea con ID: ")
// 	fmt.Println(idTask)
// 	var taskFounded = repo.GetSingleDB(idTask)
// 	taskFounded.Status = statusTask
// 	taskFounded.UpdatedAt = time.Now().String()
// 	spew.Dump(taskFounded)
// 	err := repo.db.Update(taskFounded)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Actualizado con éxito")
// }
