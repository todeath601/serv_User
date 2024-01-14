package main

import (
	"fmt"
	"log"
	"page/database"
	"page/handlers"
	"page/service"

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

func CreateTableUser() {
	storage := database.NewStorage()
	err := storage.CreateSchema()
	if err != nil {
		log.Fatal(err)
	}
}

func CreateDB() {
	storage := database.NewPostgresStorage()
	defer storage.Close()
}

func InitLogger() {
	service.Init()
	p := database.NewPostgresStorage()
	users := p.Read()
	fmt.Println(users)
}

func main() {

	InitLogger()
	CreateDB()
	CreateTableUser()
	gin.SetMode(gin.ReleaseMode)
	router := getRouter()

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
