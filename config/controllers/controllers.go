package controllers

import (
	"net/http"

	"github.com/eron97/crud_nov.git/config/database"
	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	db := database.GetDB(c)

	// Chama a função QueryAllTasks do pacote database
	tasks, err := database.QueryAllTasks(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar tasks"})
		return
	}

	// Cria uma slice para armazenar os nomes das tasks
	var taskNames []string

	// Itera sobre as tasks e extrai o nome de cada uma
	for _, task := range tasks {
		taskNames = append(taskNames, task.Task_Name)
	}

	// Retorna os nomes das tasks como resposta
	c.JSON(http.StatusOK, gin.H{"tasks": taskNames})
}

// GetTaskByName retorna os detalhes de uma tarefa pelo nome
func GetTaskByName(c *gin.Context) {
	// Obtém o ponteiro do banco de dados do contexto Gin
	db := database.GetDB(c)

	// Obtém o nome da task dos parâmetros da solicitação
	taskName := c.Param("task_name")

	// Chama a função GetTaskByName do pacote database
	task, err := database.GetTaskByName(db, taskName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar a task"})
		return
	}

	// Retorna os detalhes da task como resposta
	c.JSON(http.StatusOK, gin.H{"task": task})
}
