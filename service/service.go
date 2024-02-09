package service

type User struct {
	ID        string `json:"id" validate:"required,min=1"`
	FirstName string `json:"first_name" validate:"required,min=2,max=16"`
	LastName  string `json:"last_name" validate:"required,min=2,max=20"`
	Age       int    `json:"age" validate:"gte=18"`
}

type Storage interface {
	CreateSchema() error
	Create(User) User
	Read() []User
	ReadOne(id string) (User, error)
	Delete(id string) error
}
