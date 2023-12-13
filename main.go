package main

import (
	"page/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/ping", handlers.Pong)
	r.GET("/user", handlers.GetUser)
	r.POST("/user", handlers.PostUser)
	r.Run(":8080")
}
