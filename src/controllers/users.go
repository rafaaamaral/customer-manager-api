package controllers

import (
	"customer-manager-api/src/database"
	"customer-manager-api/src/models"
	"customer-manager-api/src/repositories"
	"customer-manager-api/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err := json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := user.BeforeCreate(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	repository := repositories.NewUsersRepository(database.DB)
	user.ID, err = repository.CreateUser(user)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting Users"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting User"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating User"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting User"))
}
