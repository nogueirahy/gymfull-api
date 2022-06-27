package main

import (
	"github.com/nogueirahy/api-gymfull/database"
	"github.com/nogueirahy/api-gymfull/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()
	server.Run()
}
