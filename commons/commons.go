package commons

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/DjonatasBorges/api-user/errors"
	"github.com/go-playground/validator/v10"
)

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("Error:", err)
	}
}

func GenerateRandomKey(length int) (string, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

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

func ValidateCPF(cpf string) bool {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	if len(cpf) != 11 {
		return false
	}

	for i := 1; i < 11; i++ {
		if cpf[i] != cpf[0] {
			break
		}
		if i == 10 {
			return false
		}
	}

	sum := 0
	for i := 0; i < 9; i++ {
		digit, err := strconv.Atoi(string(cpf[i]))
		if err != nil {
			return false
		}
		sum += digit * (10 - i)
	}
	remainder := sum % 11
	digit1 := 0
	if remainder >= 2 {
		digit1 = 11 - remainder
	}

	if strconv.Itoa(digit1) != string(cpf[9]) {
		return false
	}

	sum = 0
	for i := 0; i < 10; i++ {
		digit, err := strconv.Atoi(string(cpf[i]))
		if err != nil {
			return false
		}
		sum += digit * (11 - i)
	}
	remainder = sum % 11
	digit2 := 0
	if remainder >= 2 {
		digit2 = 11 - remainder
	}

	if strconv.Itoa(digit2) != string(cpf[10]) {
		return false
	}

	return true
}
