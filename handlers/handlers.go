package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

type error struct {
	Error string `json:"error`
}

var users = []user{
	{ID: "1", FirstName: "Vladimir", LastName: "Sakhonchyk", Age: 24},
	{ID: "2", FirstName: "Nikita", LastName: "Samokhvalov", Age: 25},
	{ID: "3", FirstName: "Alina", LastName: "Makarenko", Age: 22},
}

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func DeleteUsersById(c *gin.Context) {
	id := c.Param("id")
	for i, a := range users {
		if a.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.IndentedJSON(http.StatusNoContent, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, error{"not found"})
}

func GetUsersById(c *gin.Context) {
	id := c.Param("id")
	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
		}
	}
	c.IndentedJSON(http.StatusNotFound, error{"not found"})
}

func PostUsers(c *gin.Context) {
	var newUser user
	err := c.BindJSON(&newUser)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, error{"bad request"})
		return
	}
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
