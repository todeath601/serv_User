package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Storage interface {
	Create() user
	Read() user
	ReadOne() (user, error)
	Delete() user
}

type user struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

type MemoryStorage struct {
	users []user
}

func (s MemoryStorage) Create(us user) user {
	s.users = append(s.users, us)
	return us

}

func (s MemoryStorage) ReadOne(id string) (user, error) {
	for _, a := range s.users {
		if a.ID == id {
			return a, nil
		}
	}
	return user{}, errors.New("not found")
}

func (s MemoryStorage) Delete(id string) error {
	for i, a := range s.users {
		if a.ID == id {
			s.users = append(s.users[:i], s.users[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")

}

func (s MemoryStorage) Read() []user {
	return s.users
}

func NewMemoryStorage() MemoryStorage {
	var users = []user{
		{ID: "1", FirstName: "Vladimir", LastName: "Sakhonchyk", Age: 24},
		{ID: "2", FirstName: "Nikita", LastName: "Samokhvalov", Age: 25},
		{ID: "3", FirstName: "Alina", LastName: "Makarenko", Age: 22},
	}
	return MemoryStorage{users: users}
}

type HttpError struct {
	Error string `json:"error"`
}

var users = []user{
	{ID: "1", FirstName: "Vladimir", LastName: "Sakhonchyk", Age: 24},
	{ID: "2", FirstName: "Nikita", LastName: "Samokhvalov", Age: 25},
	{ID: "3", FirstName: "Alina", LastName: "Makarenko", Age: 22},
}

func GetUsers(c *gin.Context) {
	storage := NewMemoryStorage()
	c.IndentedJSON(http.StatusOK, storage.Read())
}

func DeleteUsersById(c *gin.Context) {
	id := c.Param("id")
	storage := NewMemoryStorage()
	err := storage.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, HttpError{"not found"})

		return
	}
	c.IndentedJSON(http.StatusNoContent, user{})
}

func GetUsersById(c *gin.Context) {
	id := c.Param("id")
	storage := NewMemoryStorage()
	user, err := storage.ReadOne(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, HttpError{"not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, user)

	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, HttpError{"not found"})
}

func PostUsers(c *gin.Context) {
	var newUser user
	err := c.BindJSON(&newUser)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, HttpError{"bad request"})
		return
	}
	storage := NewMemoryStorage()
	storage.Create(newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

// Neki4 4elebosik, esli yvidel, to ya sdelay zavtra(segodny)14122023
