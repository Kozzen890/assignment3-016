package main

import (
	"github.com/Kozzen890/assignment3-016/controllers"
	"github.com/Kozzen890/assignment3-016/database"
	"github.com/Kozzen890/assignment3-016/routers"
)

func main() {
	database.DbInit()

	go controllers.UpdateData()

	route := routers.InitRoutes()
	route.Run(":8080")
}