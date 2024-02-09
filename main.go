package main

import (
	"log"
	"page/database"
	"page/router"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

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
	r := router.GetRouter(logger)
	r.Use(router.Autorize(logger))
	CreateDB(logger)
	CreateTableUser(logger)
	gin.SetMode(gin.ReleaseMode)
	router.GetRouter(logger)
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка при запуске сервера:", err)
	}

}
