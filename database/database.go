package database

import (
	"database/sql"
	"errors"
	"page/service"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type MemoryStorage struct {
	users []service.User
}

func (s *MemoryStorage) Create(us service.User) service.User {
	service.Logger.WithFields(logrus.Fields{
		"action":    "create_user",
		"user_data": us,
	}).Info("Creating user in memory storage")

	createUser := append(s.users, us)
	s.users = createUser

	return us
}

func (s *MemoryStorage) Read() []service.User {
	service.Logger.WithFields(logrus.Fields{
		"action": "read_users",
	}).Info("Reading users from memory storage")

	return s.users
}

func (s *MemoryStorage) ReadOne(id string) (service.User, error) {
	for _, a := range s.users {
		if a.ID == id {
			service.Logger.WithFields(logrus.Fields{
				"action": "read_one_user",
				"id":     id,
			}).Info("User found in memory storage")

			return a, nil
		}
	}

	service.Logger.WithFields(logrus.Fields{
		"action": "read_one_user",
		"id":     id,
	}).Error("User not found in memory storage")

	return service.User{}, errors.New("not found")
}

func (s *MemoryStorage) Delete(id string) error {
	for i, a := range s.users {
		if a.ID == id {
			deleteUser := append(s.users[:i], s.users[i+1:]...)
			s.users = deleteUser

			service.Logger.WithFields(logrus.Fields{
				"action": "delete_user",
				"id":     id,
			}).Info("User deleted successfully")

			return nil
		}
	}

	service.Logger.WithFields(logrus.Fields{
		"action": "delete_user",
		"id":     id,
	}).Error("User not found")

	return errors.New("not found")
}

func (s *MemoryStorage) Close() {
	if s != nil {
		s.Close()
	}
}

// func NewMemoryStorage(storage service.Storage) *MemoryStorage {
// 	users := []service.User{
// 		{ID: "6", FirstName: "Vladimir", LastName: "Sakhonchyk", Age: 24},
// 		{ID: "7", FirstName: "Nikita", LastName: "Samokhvalov", Age: 25},
// 		{ID: "8", FirstName: "Alina", LastName: "Makarenko", Age: 22},
// 	}
// 	for _, user := range users {
// 		storage.Create(user)
// 	}
// 	return &MemoryStorage{users: users}
// }

type PostgresStorage struct {
	db *sql.DB
}

func (p *PostgresStorage) CreateSchema() error {
	_, err := p.db.Exec("SELECT 1 FROM users LIMIT 1;")
	if err != nil {
		_, err = p.db.Exec("CREATE TABLE IF NOT EXISTS users (ID SERIAL PRIMARY KEY, FirstName VARCHAR(50), LastName VARCHAR(50), Age integer);")
		if err != nil {
			service.Logger.WithError(err).Error("Error creating the users table schema")
			return err
		}
		service.Logger.Info("Users table schema created successfully")
	} else {
		service.Logger.Info("Users table already exists")
	}
	return nil
}

func NewPostgresStorage() *PostgresStorage {
	connStr := "port=5432 host=localhost user=user password=admin dbname=db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logrus.Fatal(err)
	}
	db.Stats()

	storage := &PostgresStorage{db: db}
	err = storage.CreateSchema()
	if err != nil {
		logrus.Fatal(err)
	}
	return storage
}

func (p *PostgresStorage) Close() {
	if p.db != nil {
		p.db.Close()
	}
}

func (p *PostgresStorage) Create(us service.User) service.User {
	service.Logger.WithFields(logrus.Fields{
		"action": "create_user",
	}).Info("Create user")
	_, err := p.db.Exec("INSERT INTO users(ID, FirstName, LastName, Age) VALUES($1, $2, $3, $4)", us.ID, us.FirstName, us.LastName, us.Age)
	if err != nil {
		service.Logger.WithError(err).Error("Error while creating user")
	}
	return us
}

func (p *PostgresStorage) ReadOne(id string) (service.User, error) {
	var user service.User
	row := p.db.QueryRow("SELECT * FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			service.Logger.WithFields(logrus.Fields{
				"action": "read_user",
				"id":     id,
			}).Info("User not found")
		} else {
			service.Logger.WithFields(logrus.Fields{
				"action": "read_user",
				"id":     id,
			}).Error("Error executing query")
		}
		return user, err
	}

	service.Logger.WithFields(logrus.Fields{
		"action": "read_user",
		"id":     id,
	}).Info("User successfully received")

	return user, nil
}

func (p *PostgresStorage) Delete(id string) error {
	_, err := p.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		service.Logger.WithError(err).WithField("userID", id).Error("Error deleting user")
		return err
	}

	service.Logger.WithField("userID", id).Info("User deleted successfully")
	return nil
}

func (p *PostgresStorage) Read() []service.User {
	var users []service.User
	rows, err := p.db.Query("SELECT * FROM users")
	if err != nil {
		service.Logger.WithError(err).Fatal("Error executing query")
	}
	defer rows.Close()

	for rows.Next() {
		var u service.User
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age)
		if err != nil {
			service.Logger.WithError(err).Fatal("Error scanning row")
		}
		users = append(users, u)
	}
	service.Logger.WithField("userCount", len(users)).Info("Read users successfully")
	return users
}

func NewStorage() service.Storage {
	return NewPostgresStorage()
}
