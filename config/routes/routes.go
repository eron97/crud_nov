package routes

import (
	"github.com/eron97/crud_nov.git/config/controllers"
	"github.com/eron97/crud_nov.git/config/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupTaskRoutes(router *gin.Engine) {
	router.GET("/tasks", controllers.GetAllTasks)
	router.GET("/tasks/:task_name", controllers.GetTaskByName)
	router.POST("/tasks", middlewares.ValidateTaskData, controllers.AddTask)
}
