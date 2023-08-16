package repositories

import (
	"database/sql"

	"github.com/DjonatasBorges/api-user/database"
	"github.com/DjonatasBorges/api-user/errors"
	"github.com/DjonatasBorges/api-user/models"
	"github.com/google/uuid"
)

func GetAllUsers(name string) (*sql.Rows, error) {
	db := database.ConnectDB()
	defer db.Close()

	var query string
	var rows *sql.Rows
	var err error

	if name != "" {
		query = "SELECT * FROM users WHERE LOWER(name) = LOWER($1) OR LOWER(lastname) = LOWER($1)"
		rows, err = db.Query(query, name)
	} else {
		query = "SELECT * FROM users"
		rows, err = db.Query(query)
	}

	return rows, err
}

func SaveUser(newUser models.User) (*models.User, error) {
	db := database.ConnectDB()
	defer db.Close()

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
	return &newUser, nil
}

func GetUserById(id uuid.UUID) *sql.Row {
	db := database.ConnectDB()
	defer db.Close()

	query := "SELECT * FROM users WHERE id = $1"
	row := db.QueryRow(query, id)

	return row
}

func DeleteUser(uuid uuid.UUID) error {
	db := database.ConnectDB()
	defer db.Close()

	query := "DELETE FROM users WHERE id = $1"
	_, err := db.Exec(query, uuid)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByEmail(email string) (*sql.Rows, error) {
	db := database.ConnectDB()
	defer db.Close()

	query := "SELECT * FROM users WHERE email = $1"
	rows, err := db.Query(query, email)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, errors.ErrUserNotFound
	}

	return rows, nil
}
