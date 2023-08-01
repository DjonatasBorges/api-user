package controllers

import (
	"encoding/json"
	"log"
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
		commons.HandleServerError(w, err)
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
		commons.WriteJSONResponse(w, http.StatusNotFound, err)
		commons.HandleServerError(w, err)
		return
	}

	commons.WriteJSONResponse(w, http.StatusOK, user)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	user, err := services.PostUser(newUser)
	if err != nil {
		commons.HandleServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	commons.WriteJSONResponse(w, http.StatusCreated, user)
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
		commons.HandleServerError(w, errNotFound)
		return
	}

	deleteErr := services.DeleteUser(id)
	if deleteErr != nil {
		commons.HandleServerError(w, deleteErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
