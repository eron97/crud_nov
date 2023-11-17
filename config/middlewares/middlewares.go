package middlewares

import (
	"log"
	"net/http"

	"github.com/eron97/crud_nov.git/config/models"
	"github.com/gin-gonic/gin"
)

func ValidateTaskData(c *gin.Context) {

	log.Println("[Middleware Val. dados chamado]")
	var taskPost models.TaskPost

	// Popula a estrutura a partir do c *gin.Context
	if err := c.ShouldBindJSON(&taskPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		log.Println("[Middleware: Erro crítico na validação dos campos JSON]")
		return
	}

	log.Printf("[Middleware: JSON validado | task_name: %s] | priority: %s]", taskPost.Task_Name, taskPost.Priority)
	c.Set("ValidateTaskData", taskPost)
	c.Next()
}
