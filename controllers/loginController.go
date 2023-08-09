package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/DjonatasBorges/api-user/commons"
	"github.com/DjonatasBorges/api-user/errors"
	"github.com/DjonatasBorges/api-user/middleware"
	"github.com/DjonatasBorges/api-user/services"
	"github.com/golang-jwt/jwt/v5"
)

type LoginCredentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials LoginCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		commons.WriteJSONResponse(w, http.StatusBadRequest, errors.ErrInvalidCredentials)
		return
	}

	if err := commons.ValidateData(credentials); err != nil {
		commons.WriteJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	user, err := services.GetUserByEmail(credentials.Email)
	if err != nil || user == nil {
		commons.WriteJSONResponse(w, http.StatusUnauthorized, errors.ErrUserNotRegistered)
		return
	}

	if user.Password != credentials.Password {
		commons.WriteJSONResponse(w, http.StatusUnauthorized, errors.ErrInvalidCredentials)
		return
	}

	expirationTime := time.Now().Add(1 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": credentials.Email,
		"exp":   expirationTime.Unix(),
	})

	tokenString, err := token.SignedString(middleware.JwtSecret)
	if err != nil {
		commons.WriteJSONResponse(w, http.StatusUnauthorized, errors.ErrInvalidTokenSignature)
		return
	}

	w.Header().Add("Authorization", "Bearer "+tokenString)
	w.WriteHeader(http.StatusOK)
}
