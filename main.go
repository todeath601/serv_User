package main

import (
	"fmt"
	"page/database"
	"page/handlers"

	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.GET("/users", handlers.GetUsers)
	router.GET("/users/:id", handlers.GetUsersById)
	router.POST("/users", handlers.PostUsers)
	router.DELETE("/users/:id", handlers.DeleteUsersById)
	return router
}

func main() {
	storage := database.NewPostgresStorage()
	defer storage.Close()
	router := getRouter()

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
