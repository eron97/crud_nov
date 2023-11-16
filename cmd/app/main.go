package main

import (
	"log"

	"github.com/eron97/crud_nov.git/config/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupTaskRoutes(r)
	log.Printf("Server On localhost:8080")
	r.Run()

}
