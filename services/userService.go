package services

import (
	"log"

	"github.com/DjonatasBorges/api-user/errors"
	"github.com/DjonatasBorges/api-user/models"
	"github.com/DjonatasBorges/api-user/repositories"
	"github.com/google/uuid"
)

func GetAllUsers() ([]models.User, error) {
	rows, err := repositories.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Lastname, &user.Cpf, &user.Email, &user.Phone); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func PostUser(newUser models.User) (*models.User, error) {
	return repositories.SaveUser(newUser)
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	row := repositories.GetUserById(id)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Lastname, &user.Cpf, &user.Email, &user.Phone)
	if err != nil {
		log.Println("Error fetching user:", err)
		return nil, errors.ErrUserNotFound.WithArgs("User")
	}

	return &user, nil
}

func DeleteUser(uuid uuid.UUID) error {
	return repositories.DeleteUser(uuid)
}
