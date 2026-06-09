package controller

import (
	"encoding/json"
	"extensao-api/models"
	"extensao-api/persistency"
	"extensao-api/repository"
	"extensao-api/responses"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var newUser models.Users
	if err = json.Unmarshal(bodyRequest, &newUser); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = newUser.Prepare("create"); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := persistency.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	repo := repository.NewUsersRepo(db)
	newUser.ID, err = repo.Create(newUser)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()
	responses.JSON(w, http.StatusCreated, newUser)
}

func FetchUser(writer http.ResponseWriter, request *http.Request) {

}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {

}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {

}
