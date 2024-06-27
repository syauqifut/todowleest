package main

import (
	"github.com/syauqifut/todowleest/controllers/todocontroller"
	"github.com/syauqifut/todowleest/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDB()

	router.GET("/", rootHandler)

	router.GET("/api/todo_api/", todocontroller.Index)
	router.POST("/api/todo_api/", todocontroller.Post)
	router.GET("/api/todo_api/:id", todocontroller.Det)
	router.PUT("/api/todo_api/:id", todocontroller.Update)
	router.DELETE("/api/todo_api/:id", todocontroller.Del)

	router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
