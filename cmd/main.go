package main

import (
	"log"
	"ozon-task/internal/app"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
