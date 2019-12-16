package main

import (
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/onyas/geekNews/router"
)

func main() {
	//port := os.Getenv("PORT")
	//
	//if port == "" {
	//	log.Fatal("$PORT must be set")
	//}

	router := router.SetupRouter()
	router.Run()
}
