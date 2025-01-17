package main

import (
	"github.com/kataras/iris/v12"
	"iris-learn/router"
	"log"
)

func main() {
	app := router.NewRouter()

	err := app.Listen(":8080", iris.WithLogLevel("debug"))
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
}
