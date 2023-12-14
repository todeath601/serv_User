package main

import (
	"page/handlers"

	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users", handlers.GetUsers)
	router.GET("/users/:id", handlers.GetUsersById)
	router.POST("/users", handlers.PostUsers)
	router.DELETE("/users/:id", handlers.DeleteUsersById)
	return router
}

func main() {
	router := getRouter()
	router.Run(":8080")
}
