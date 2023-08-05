package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/DjonatasBorges/api-user/commons"
	"github.com/DjonatasBorges/api-user/errors"
	"github.com/DjonatasBorges/api-user/models"
	"github.com/DjonatasBorges/api-user/services"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	commons.WriteJSONResponse(w, http.StatusOK, users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		errorResponse := errors.ErrInvalidUUID
		commons.WriteJSONResponse(w, http.StatusBadRequest, errorResponse)
		return
	}

	user, err := services.GetUserById(id)
	if err != nil {
		commons.WriteJSONResponse(w, http.StatusNotFound, errors.ErrUserNotFound.WithArgs("User"))
		return
	}

	commons.WriteJSONResponse(w, http.StatusOK, user)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		commons.WriteJSONResponse(w, http.StatusBadRequest, errors.ErrInvalidBodyRequest)
		return
	}

	if err := commons.ValidateData(newUser); err != nil {
		commons.WriteJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	cpfIsValid := commons.ValidateCPF(newUser.Cpf)
	if !cpfIsValid {
		commons.WriteJSONResponse(w, http.StatusBadRequest, errors.ErrInvalidCPF)
		return
	}

	user, errUser := services.GetUserByEmail(newUser.Email)
	if user.Email == newUser.Email || user.Cpf == newUser.Cpf {
		commons.WriteJSONResponse(w, http.StatusBadRequest, errors.ErrUserAlreadyExists)
	}
	if errUser != nil {
		commons.WriteJSONResponse(w, http.StatusInternalServerError, errors.ErrInternalServer)
	}

	user, errorPost := services.PostUser(newUser)
	if errorPost != nil {
		commons.WriteJSONResponse(w, http.StatusBadRequest, errorPost)
		return
	}

	w.Header().Set("userId", user.ID.String())
	w.WriteHeader(http.StatusCreated)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		errorResponse := errors.ErrInvalidUUID
		commons.WriteJSONResponse(w, http.StatusBadRequest, errorResponse)
		return
	}

	_, errNotFound := services.GetUserById(id)
	if errNotFound != nil {
		commons.WriteJSONResponse(w, http.StatusNotFound, errNotFound)
		return
	}

	deleteErr := services.DeleteUser(id)
	if deleteErr != nil {
		commons.WriteJSONResponse(w, http.StatusInternalServerError, errors.ErrInternalServer)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
