package repositories

import (
	"database/sql"

	"github.com/DjonatasBorges/api-user/database"
	"github.com/DjonatasBorges/api-user/models"
	"github.com/google/uuid"
)

func GetAllUsers() (*sql.Rows, error) {
	db := database.ConnectDB()

	query := "SELECT * FROM users"
	rows, err := db.Query(query)

	db.Close()
	return rows, err
}

func SaveUser(newUser models.User) (*models.User, error) {
	db := database.ConnectDB()

	query := `
		INSERT INTO users (name, lastname, cpf, email, phone, password)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	var id string
	err := db.QueryRow(query, newUser.Name, newUser.Lastname, newUser.Cpf, newUser.Email, newUser.Phone, newUser.Password).Scan(&id)
	if err != nil {
		return nil, err
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	newUser.ID = uuid

	defer db.Close()

	return &newUser, nil
}

func GetUserById(id uuid.UUID) *sql.Row {
	db := database.ConnectDB()

	query := "SELECT * FROM users WHERE id = $1"
	row := db.QueryRow(query, id)

	db.Close()
	return row
}

func DeleteUser(uuid uuid.UUID) error {
	db := database.ConnectDB()

	query := "DELETE FROM users WHERE id = $1"
	_, err := db.Exec(query, uuid)
	if err != nil {
		return err
	}

	db.Close()

	return nil
}

func GetUserByEmail(email string) (*sql.Rows, error) {
	db := database.ConnectDB()

	query := "SELECT * FROM users WHERE email = $1"
	rows, err := db.Query(query, email)
	if err != nil {
		return nil, err
	}

	db.Close()
	return rows, nil
}
