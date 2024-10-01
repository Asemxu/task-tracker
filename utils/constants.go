package utils

import "fmt"

//Options selecteds with state
const ADDTASK = "add"
const UPDATETASK = "update"
const DELETETASK = "delete"
const MARKPROGRESSTASK = "mark-in-progress"
const MARKDONETASK = "mark-done"
const LISTTASKS = "list"
const LISTTASKSDONE = "done"
const LISTTASKSTODO = "todo"
const LISTTASKSINPROGRESS = "in-progress"

//MESSAGES
const DONTHAVEOPTION = "No ingreso ninguna opción"
const OPTIONNOTVALID = "La opción ingresada no es válida"
const TOLESSARGUMENTS = "No se han ingresado todos los argumentos necesarios"

var GETOPTIONSELECTED = func(option string) {
	fmt.Println("La opción seleccionada fue " + option)
}

//DEFAULT VALUES
const DESCRIPTIONDEFAULT = "Esta es información por default"
const EMPTYVALUE = ""
const CANTNOTUPDATE = "No se pudo actualizar la tarea"
const ACCEPTCANTARGUMENTSUPDATE = 3
const ACCEPTCANTARGUMENTSDELETE = 2
const ACCEPTCANTARGUMENTSMARKTASK = 2
const ACCEPTCANTARGUMENTSLIST = 1
const ACCEPTCANTARGUMENTSLISTQUERY = 2

const MARKCREATED = 0
const MARKPROGRESS = 1
const MARKDONE = 2
const LISTDONE = "0"
const LISTTODO = "1"
const LISTPROGRESS = "2"

//FIELDNAMES
const DESCRIPTIONFIELD = "Description"
const UPDATEDATFIELD = "UpdatedAt"
const STATUSFIELD = "Status"
const STATUSFIELDJSON = "status"

const ISQUERY = true
const NOTQUERY = false
