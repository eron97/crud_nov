package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var tasks = [3]int{1, 2, 3}

func GetAllTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}
