package main

import (
	"fmt"
	"log"
	"page/database"
	"page/handlers"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func getRouter(logger *logrus.Logger) *gin.Engine {
	handler := handlers.NewHandler(logger)
	router := gin.Default()
	router.GET("/users", handler.GetUsers)
	router.GET("/users/:id", handler.GetUsersById)
	router.POST("/users", handler.PostUsers)
	router.DELETE("/users/:id", handler.DeleteUsersById)
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
	router := getRouter(logger)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
