package service

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

type Storage interface {
	Create(User) User
	Read() []User
	ReadOne(id string) (User, error)
	Delete(id string) error
}
