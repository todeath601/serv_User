package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       string `json:"age"`
}

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetUser(c *gin.Context) {

	user := User{"1", "Alex", "Shveden", "23"}
	c.JSON(200, gin.H{
		"user": user,
	})
}

func PostUser(c *gin.Context) {

	user := User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errror": "BadRequest"})
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}
