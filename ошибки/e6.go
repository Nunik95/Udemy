package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, errr := os.Create("log.txt")
	if errr != nil {
		fmt.Println(errr)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("Nunka is the best")
	f2, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("error happened", err)
	}
	defer f2.Close()
}
