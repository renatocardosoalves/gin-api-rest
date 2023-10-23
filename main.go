package main

import (
	"github.com/renatocardosoalves/api-go-gin/database"
	"github.com/renatocardosoalves/api-go-gin/router"
)

func main() {
	database.ConnectDatabase()
	router.HandleRequest()
}
