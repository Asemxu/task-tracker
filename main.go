package main

import (
	"fmt"
	"os"
	"task-tracker/utils"
)

func main() {
	arguments := os.Args[1:]
	fmt.Println(arguments)

	if utils.IsValidArguments(arguments) {
		fmt.Println("La opcion seleccionada fue " + arguments[0])
	} else {
		fmt.Println("No ingreso ninguna opción")
	}
	os.Exit(1)
}
