package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/DjonatasBorges/api-user/commons"
	"github.com/DjonatasBorges/api-user/errors"
	"github.com/DjonatasBorges/api-user/services"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key")

type LoginCredentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return token, nil
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

	expirationTime := time.Now().Add(24 * time.Hour) // Definindo expiração em 60 segundos

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": credentials.Email,
		"exp":   expirationTime.Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		commons.WriteJSONResponse(w, http.StatusUnauthorized, errors.ErrInvalidTokenSignature)
		return
	}

	w.Header().Add("Authorization", "Bearer "+tokenString)
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {

	commons.WriteJSONResponse(w, http.StatusUnauthorized, errors.ErrMissingAuthToken)
}

func ExtractToken(r *http.Request) (string, error) {
	bearerToken := r.Header.Get("Authorization")
	if len(bearerToken) > 7 && strings.ToUpper(bearerToken[0:7]) == "BEARER " {
		return bearerToken[7:], nil
	}

	return "", nil
}
