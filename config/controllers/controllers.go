package controllers

import (
	"log"
	"net/http"

	"github.com/eron97/crud_nov.git/config/database"
	"github.com/eron97/crud_nov.git/config/models"
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
	task, err := database.QueryByName(db, taskName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar a task"})
		return
	}

	// Retorna os detalhes da task como resposta
	c.JSON(http.StatusOK, gin.H{"task": task})
}

func AddTask(c *gin.Context) {

	log.Println("[Controllers: Func AddTask iniciada]")
	db := database.GetDB(c)

	ValidateTaskData, exists := c.Get("ValidateTaskData")
	if !exists {
		log.Println("[Controllers: Erro ao recuperar *gin.Context]")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno ao processar"})
		return
	}

	taskPost, ok := ValidateTaskData.(models.TaskPost)
	if !ok {
		log.Println("[Controllers: Erro ao criar struct models.TaskPost a partir de *gin.Context]")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao converter os dados"})
		return
	}

	log.Println("[Controllers: chamando database.QueryAddTask]")

	// Chama a função AddTask do pacote database
	err := database.QueryAddTask(db, &taskPost)
	if err != nil {
		log.Println("[Controllers: Erro ao executar a função database.QueryAddTask]")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar a tarefa"})
		return
	}

	log.Println("[Controllers: Nova tarefa registrada com sucesso]")
	c.JSON(http.StatusCreated, gin.H{"message": "Tarefa adicionada com sucesso"})
}
