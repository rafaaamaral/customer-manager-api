package controllers

import (
	"customer-manager-api/src/database"
	"customer-manager-api/src/models"
	"customer-manager-api/src/repositories"
	"customer-manager-api/src/responses"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if err := user.PrepareToSave(false); err != nil {
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
	name := strings.ToLower(r.URL.Query().Get("name"))

	repository := repositories.NewUsersRepository(database.DB)
	users, err := repository.GetUsers(name)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	repository := repositories.NewUsersRepository(database.DB)
	user, err := repository.GetUserById(uint(userID))

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
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

	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	user.ID = uint(userID)
	if err := user.PrepareToSave(true); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	repository := repositories.NewUsersRepository(database.DB)

	err = repository.UpdateUser(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	repository := repositories.NewUsersRepository(database.DB)
	err = repository.DeleteUser(uint(userID))

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
