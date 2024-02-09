package router

import (
	"net/http"
	"page/handlers"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

)

func GetRouter(logger *logrus.Logger) *gin.Engine {
	handler := handlers.NewHandler(logger)
	router := gin.Default()
	router.Use(Autorize(logger))
	router.GET("/users", BasicAuth(), handler.GetUsers)
	router.GET("/users/:id", BasicAuth(), handler.GetUsersById)
	router.POST("/users", BasicAuth(), handler.PostUsers)
	router.DELETE("/users/:id", BasicAuth(), handler.DeleteUsersById)

	return router
}

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"username": "password",
	})

}

func Autorize(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get(gin.AuthUserKey)
		if !exists {
			logger.WithFields(logrus.Fields{"status": http.StatusUnauthorized, "error": "Unauthorized"}).Warn("Access unauthorized")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		if user.(string) != "admin" {
			logger.WithFields(logrus.Fields{"status": http.StatusForbidden, "error": "Forbidden"}).Warn("Access forbidden")
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}
		c.Set("admin", user.(string))

		c.Next()
	}
}
