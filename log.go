package main

import (
	"log"
	"os"
)

func startLogger() *os.File {
	file, _ := os.Create("nyan.log")
	log.SetOutput(file)
	log.Println("Start Nyan Server")

	return file
}
