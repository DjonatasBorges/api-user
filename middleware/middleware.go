package middleware

import (
	"net/http"
	"strings"

	"github.com/DjonatasBorges/api-user/commons"
	"github.com/DjonatasBorges/api-user/errors"
	"github.com/golang-jwt/jwt/v5"
)

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}

var JwtSecret = []byte("your-secret-key")

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := ExtractToken(r)
		if err != nil {
			commons.WriteJSONResponse(w, http.StatusBadRequest, errors.ErrMissingAuthToken)
			return
		}

		token, err := ParseToken(tokenString)
		if err != nil || !token.Valid {
			commons.WriteJSONResponse(w, http.StatusUnauthorized, errors.ErrInvalidAuthHeader)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return token, nil
}

func ExtractToken(r *http.Request) (string, error) {
	bearerToken := r.Header.Get("Authorization")
	if len(bearerToken) > 7 && strings.ToUpper(bearerToken[0:7]) == "BEARER " {
		return bearerToken[7:], nil
	}

	return "", nil
}
