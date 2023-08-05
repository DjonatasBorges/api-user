package commons

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/DjonatasBorges/api-user/errors"
	"github.com/go-playground/validator/v10"
)

func HandleServerError(w http.ResponseWriter, err error) {
	log.Println("Error:", err)
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		HandleServerError(w, err)
	}
}

// AuthErrorResponse representa a estrutura da resposta de erro de autenticação.
type AuthErrorResponse struct {
	Error string `json:"error"`
}

// HandleAuthError envia uma resposta JSON para o cliente indicando um erro de autenticação.
func HandleAuthError(w http.ResponseWriter, err error) {
	response := AuthErrorResponse{
		Error: err.Error(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(response)
}

func GenerateRandomKey(length int) (string, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	// Encode em base64 para tornar a chave mais amigável para uso em strings.
	encodedKey := base64.RawStdEncoding.EncodeToString(key)
	return encodedKey, nil
}

func ValidateData(data interface{}) error {
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		var invalidFields []string
		for _, e := range err.(validator.ValidationErrors) {
			invalidFields = append(invalidFields, e.Field())
		}
		errMsg := strings.Join(invalidFields, ", ")

		appErr := errors.ErrFieldCannotBeNull.WithArgs(errMsg)

		return appErr
	}
	return nil
}
