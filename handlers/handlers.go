package handlers

import (
	"net/http"
	"page/database"
	"page/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type HttpError struct {
	Error string `json:"error"`
}

func GetUsers(c *gin.Context) {
	logger := logrus.New()
	storage := database.NewStorage(logger)
	logger.WithFields(logrus.Fields{
		"endpoint": "/users",
		"method":   "GET",
	}).Info("Handling GET request for users")
	users := storage.Read()
	logger.WithFields(logrus.Fields{
		"users_count": len(users),
	}).Info("Retrieved users from the storage")
	c.IndentedJSON(http.StatusOK, users)
}

// func DeleteUsersById(c *gin.Context) {
// 	logger := logrus.New()
// 	storage := database.NewStorage(logger)
// 	id := c.Param("id")
// 	logger.WithFields(logrus.Fields{
// 		"endpoint": "/users/:id",
// 		"method":   "DELETE",
// 		"id":       id,
// 	}).Info("Handling DELETE request for user by ID")
// 	err := storage.Delete(id)
// 	if err != nil {
// 		logger.WithError(err).Error("Error deleting user by ID")
// 		c.IndentedJSON(http.StatusNotFound, HttpError{Error: "not found"})
// 		return
// 	}
// 	logger.WithFields(logrus.Fields{
// 		"id": id,
// 	}).Info("User deleted successfully")
// 	c.IndentedJSON(http.StatusNoContent, nil)
// }

func GetUsersById(c *gin.Context) {
	logger := logrus.New()
	storage := database.NewStorage(logger)
	id := c.Param("id")
	logger.WithFields(logrus.Fields{
		"endpoint": "/users/:id",
		"method":   "GETBYID",
		"id":       id,
	}).Info("Handling GET request for user by ID")
	user, err := storage.ReadOne(id)
	if err != nil {
		logger.WithError(err).Error("Error receiving user by ID")
		c.IndentedJSON(http.StatusNotFound, HttpError{Error: "not found"})
		return
	}
	logger.WithFields(logrus.Fields{
		"id": id,
	}).Info("User successfully received")
	c.IndentedJSON(http.StatusOK, user)
}

func PostUsers(c *gin.Context) {

	logger := logrus.New()
	storage := database.NewStorage(logger)
	var newUser service.User
	logger.WithFields(logrus.Fields{
		"endpoint": "/users",
		"method":   "POST",
	}).Info("Handling POST request for creating a new user")
	logger.WithFields(logrus.Fields{
		"user_data": newUser,
	}).Debug("Received user data")
	if err := c.BindJSON(&newUser); err != nil {
		logger.WithError(err).Error("Error binding JSON data")
		c.IndentedJSON(http.StatusBadRequest, HttpError{Error: "bad request"})
		return
	}
	storage.Create(newUser)
	logger.WithFields(logrus.Fields{
		"user_id": newUser.ID,
	}).Info("User created successfully")
	c.IndentedJSON(http.StatusCreated, newUser)
}

// Neki4 4elebosik, esli yvidel, to ya sdelay zavtra(segodny)14122023
