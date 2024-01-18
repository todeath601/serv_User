package service

import "github.com/sirupsen/logrus"

var Logger = logrus.New()

func Init() {

	Logger.Formatter = &logrus.TextFormatter{}
	Logger.Level = logrus.DebugLevel

}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

type Storage interface {
	CreateSchema() error
	Create(User) User
	Read() []User
	ReadOne(id string) (User, error)
	Delete(id string) error
}