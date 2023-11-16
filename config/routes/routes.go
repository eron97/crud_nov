package routes

import (
	"github.com/eron97/crud_nov.git/config/controllers"
	"github.com/gin-gonic/gin"
)

func SetupTaskRoutes(router *gin.Engine) {
	// router.GET("/tasks", controllers.GetAllTasks)
	router.GET("/tasks", controllers.GetAllTasks)
}
