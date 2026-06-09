package controller

import (
	"encoding/json"
	"extensao-api/auth"
	"extensao-api/models"
	"extensao-api/persistency"
	"extensao-api/repository"
	"extensao-api/responses"
	"extensao-api/security"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, 400, err)
		return
	}
	var user models.Users
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Err(w, 400, err)
		return
	}
	db, err := persistency.Connect()
	if err != nil {
		responses.Err(w, 500, err)
		return
	}
	defer db.Close()
	repo := repository.NewUsersRepo(db)
	userDB, err := repo.FetchByEmail(user.Email)
	if err != nil {
		responses.Err(w, 401, err)
		return
	}
	if err = security.ValidatePassword(
		userDB.Password,
		user.Password,
	); err != nil {
		responses.Err(w, 401, err)
		return
	}
	token, err := auth.GenerateToken(
		int64(userDB.ID),
	)
	if err != nil {
		responses.Err(w, 500, err)
		return
	}
	responses.JSON(
		w,
		http.StatusOK,
		map[string]string{
			"token": token,
		},
	)
}
