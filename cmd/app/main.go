package main

import (
	"log"

	"github.com/eron97/crud_nov.git/config/database"
	"github.com/eron97/crud_nov.git/config/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa o banco de dados
	db := database.InitDB()
	defer db.Close()
	log.Println("[Conexão com database ok]")

	r := gin.Default()

	// Armazena o ponteiro do banco de dados no contexto Gin
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	routes.SetupTaskRoutes(r)
	log.Println("[Rotas ok]")

	r.Run()

}
