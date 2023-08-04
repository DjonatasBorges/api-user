package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/DjonatasBorges/api-user/commons"
	"github.com/DjonatasBorges/api-user/services"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key")

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
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
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	user, err := services.GetUserByEmail(credentials.Email)
	if err != nil {
		http.Error(w, "Usuário não cadastrado", http.StatusUnauthorized)
		return
	}

	if user.Password != credentials.Password {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": credentials.Email,
		"exp":   expirationTime.Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response := LoginResponse{Token: tokenString}
	commons.WriteJSONResponse(w, http.StatusOK, response)
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Esta é uma rota protegida. Somente usuários autenticados podem acessá-la."))
}

func ExtractToken(r *http.Request) (string, error) {
	bearerToken := r.Header.Get("Authorization")
	if len(bearerToken) > 7 && strings.ToUpper(bearerToken[0:7]) == "BEARER " {
		return bearerToken[7:], nil
	}

	return "", errors.New("Token not found")
}
