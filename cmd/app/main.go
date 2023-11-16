package main

import (
	"log"

	"github.com/eron97/crud_nov.git/config/database"
	"github.com/eron97/crud_nov.git/config/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db := database.InitDB()
	defer db.Close()
	log.Println("Database On!")

	r := gin.Default()
	routes.SetupTaskRoutes(r)
	log.Println("Routes OK!")

	r.Run()

}
