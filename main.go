package main

import (
	"fmt"
	"log"
	"page/database"
	"page/handlers"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func getRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users", handlers.GetUsers)
	router.GET("/users/:id", handlers.GetUsersById)
	router.POST("/users", handlers.PostUsers)
	// router.DELETE("/users/:id", handlers.DeleteUsersById)
	return router
}

func CreateTableUser(logger *logrus.Logger) {
	storage := database.NewStorage(logger)
	err := storage.CreateSchema()
	if err != nil {
		log.Fatal(err)
	}
}
func CreateDB(logger *logrus.Logger) {
	storage := database.NewPostgresStorage(logger)
	defer storage.Close()
}

func main() {
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{}
	logger.Level = logrus.DebugLevel
	logger.Out = logrus.StandardLogger().Out
	CreateDB(logger)
	CreateTableUser(logger)
	gin.SetMode(gin.ReleaseMode)
	router := getRouter()

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
