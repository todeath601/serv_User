package handlers

import (
	"net/http"
	"page/database"
	"page/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type HttpError struct {
	Error string `json:"error"`
}

type Handler struct {
	logger    *logrus.Logger
	validator *validator.Validate
}

func NewHandler(logger *logrus.Logger) *Handler {
	return &Handler{
		logger:    logger,
		validator: validator.New(),
	}
}

func (h *Handler) ValidateStruct(s interface{}) error {
	return h.validator.Struct(s)
}

func (h *Handler) GetUsers(c *gin.Context) {
	storage := database.NewStorage(h.logger)
	h.logger.WithFields(logrus.Fields{
		"endpoint": "/users",
		"method":   "GET",
	}).Info("Handling GET request for users")
	users := storage.Read()
	h.logger.WithFields(logrus.Fields{
		"users_count": len(users),
	}).Info("Retrieved users from the storage")
	c.IndentedJSON(http.StatusOK, users)
}

func (h *Handler) DeleteUsersById(c *gin.Context) {
	storage := database.NewStorage(h.logger)
	id := c.Param("id")
	h.logger.WithFields(logrus.Fields{
		"endpoint": "/users/:id",
		"method":   "DELETE",
		"id":       id,
	}).Info("Handling DELETE request for user by ID")
	err := storage.Delete(id)
	if err != nil {
		h.logger.WithError(err).Error("Error deleting user by ID")
		c.IndentedJSON(http.StatusNotFound, HttpError{Error: "not found"})
		return
	}
	h.logger.WithFields(logrus.Fields{
		"id": id,
	}).Info("User deleted successfully")
	c.IndentedJSON(http.StatusNoContent, nil)
}

func (h *Handler) GetUsersById(c *gin.Context) {

	storage := database.NewStorage(h.logger)
	id := c.Param("id")
	logrus.WithFields(logrus.Fields{
		"endpoint": "/users/:id",
		"method":   "GETBYID",
		"id":       id,
	}).Info("Handling GET request for user by ID")
	user, err := storage.ReadOne(id)
	if err != nil {
		h.logger.WithError(err).Error("Error receiving user by ID")
		c.IndentedJSON(http.StatusNotFound, HttpError{Error: "not found"})
		return
	}
	h.logger.WithFields(logrus.Fields{
		"id": id,
	}).Info("User successfully received")
	c.IndentedJSON(http.StatusOK, user)
}

func (h *Handler) PostUsers(c *gin.Context) {
	storage := database.NewStorage(h.logger)
	var newUser service.User
	h.logger.WithFields(logrus.Fields{
		"endpoint": "/users",
		"method":   "POST",
	}).Info("Handling POST request for creating a new user")
	h.logger.WithFields(logrus.Fields{
		"user_data": newUser,
	}).Debug("Received user data")
	if err := c.BindJSON(&newUser); err != nil {
		h.logger.WithError(err).Error("Error binding JSON data")
		c.IndentedJSON(http.StatusBadRequest, HttpError{Error: "bad request"})
		return
	}
	if err := h.ValidateStruct(newUser); err != nil {
		h.logger.WithError(err).Error("Validate error")
		c.IndentedJSON(http.StatusBadRequest, HttpError{Error: "validation error"})
		return
	}
	storage.Create(newUser)
	h.logger.WithFields(logrus.Fields{
		"user_id": newUser.ID,
	}).Info("User created successfully")
	c.IndentedJSON(http.StatusCreated, newUser)
}

// Neki4 4elebosik, esli yvidel, to ya sdelay zavtra(segodny)14122023
