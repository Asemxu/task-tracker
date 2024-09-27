package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("readme.md")
	arguments := os.Args[1:]
	fmt.Println(arguments)
	if err != nil {
		log.Fatal(err)
	}
	f, err := file.Stat()

	if err != nil {
		log.Fatal(err)
	}
	size := f.Size()

	data := make([]byte, size)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
