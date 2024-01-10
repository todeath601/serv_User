package handlers

import (
	"net/http"
	"page/database"
	"page/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type MemoryStorage struct {
	users []service.User
}

type HttpError struct {
	Error string `json:"error"`
}

func GetUsers(c *gin.Context) {
	storage := database.NewStorage()
	c.IndentedJSON(http.StatusOK, storage.Read())
}

func DeleteUsersById(c *gin.Context) {
	storage := database.NewStorage()
	id := c.Param("id")
	err := storage.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, HttpError{Error: "not found"})
		return
	}
	c.IndentedJSON(http.StatusNoContent, nil)
}

func GetUsersById(c *gin.Context) {
	storage := database.NewStorage()
	id := c.Param("id")
	user, err := storage.ReadOne(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, HttpError{Error: "not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func PostUsers(c *gin.Context) {
	storage := database.NewStorage()
	var newUser service.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, HttpError{Error: "bad request"})
		return
	}
	storage.Create(newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

// Neki4 4elebosik, esli yvidel, to ya sdelay zavtra(segodny)14122023
