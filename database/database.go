package database

import (
	"database/sql"
	"errors"
	"log"
	"page/service"

	_ "github.com/lib/pq"
)

type MemoryStorage struct {
	users []service.User
}

func (s *MemoryStorage) Create(us service.User) service.User {
	createUser := append(s.users, us)
	s.users = createUser
	return us
}

func (s *MemoryStorage) Read() []service.User {
	return s.users
}

func (s *MemoryStorage) ReadOne(id string) (service.User, error) {
	for _, a := range s.users {
		if a.ID == id {
			return a, nil
		}
	}
	return service.User{}, errors.New("not found")
}

func (s *MemoryStorage) Delete(id string) error {
	for i, a := range s.users {
		if a.ID == id {
			deleteUser := append(s.users[:i], s.users[i+1:]...)
			s.users = deleteUser
			return nil
		}
	}
	return errors.New("not found")
}

func NewMemoryStorage() *MemoryStorage {
	var users = []service.User{
		{ID: "1", FirstName: "Vladimir", LastName: "Sakhonchyk", Age: 24},
		{ID: "2", FirstName: "Nikita", LastName: "Samokhvalov", Age: 25},
		{ID: "3", FirstName: "Alina", LastName: "Makarenko", Age: 22},
	}
	return &MemoryStorage{users: users}
}

type PostgresStorage struct {
	db *sql.DB
}

func (p *PostgresStorage) CreateSchema() error {
	_, err := p.db.Exec("CREATE TABLE IF NOT EXISTS users (ID SERIAL PRIMARY KEY, FirstName VARCHAR(50), LastName VARCHAR(50), Age integer)")
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func NewPostgresStorage() *PostgresStorage {
	connStr := "port=5432 host=localhost user=user password=admin dbname=db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	storage := &PostgresStorage{db: db}
	err = storage.CreateSchema()
	if err != nil {
		log.Fatal(err)
	}
	return storage
}

func (p *PostgresStorage) Close() {
	if p.db != nil {
		p.db.Close()
	}
}

func (p *PostgresStorage) Create(us service.User) service.User {
	_, err := p.db.Exec("INSERT INTO users(ID, FirstName, LastName, Age) VALUES($1, $2, $3, $4)", us.ID, us.FirstName, us.LastName, us.Age)
	if err != nil {
		log.Fatal(err)
	}
	return us
}

func (p *PostgresStorage) ReadOne(id string) (service.User, error) {
	var user service.User
	row := p.db.QueryRow("SELECT * FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("not found")
		}
		return user, err
	}
	return user, nil
}

func (p *PostgresStorage) Delete(id string) error {
	_, err := p.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresStorage) Read() []service.User {
	var users []service.User
	rows, err := p.db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var u service.User
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	return users
}

func NewStorage() service.Storage {
	return NewPostgresStorage()
}
