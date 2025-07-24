package controllers

import (
	"customer-manager-api/src/database"
	"customer-manager-api/src/models"
	"customer-manager-api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(w, err, http.StatusBadRequest)
		return
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	repository := repositories.NewUsersRepository(db)
	userID, err := repository.CreateUser(user)

	if err != nil {
		log.Fatal(w, "Failed to create user:", err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User created with ID: %d", userID)))
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
