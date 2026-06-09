package controller

import (
	"context"
	"encoding/json"
	"extensao-api/models"
	"extensao-api/persistency"
	"extensao-api/repository"
	"extensao-api/responses"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, 500, err)
		return
	}
	query := strings.TrimSpace(string(b))
	googleURL := "https://www.googleapis.com/books/v1/volumes?q=" +
		url.QueryEscape(query)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()
	gReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		googleURL,
		nil,
	)
	if err != nil {
		responses.Err(w, 500, err)
		return
	}
	resp, err := http.DefaultClient.Do(gReq)
	if err != nil {
		responses.Err(w, 500, err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		responses.Err(w, 500, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		responses.Err(w, 400, err)
		return
	}
	db, err := persistency.Connect()
	if err != nil {
		responses.Err(w, 500, err)
		return
	}
	defer db.Close()
	repo := repository.NewBooksRepo(db)
	err = repo.Create(book)
	if err != nil {
		responses.Err(w, 500, err)
		return
	}
	responses.JSON(
		w,
		http.StatusCreated,
		book,
	)
}

func FetchBooks(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("userId")
	userID, _ := strconv.Atoi(userIDStr)
	db, err := persistency.Connect()
	if err != nil {
		responses.Err(w, 500, err)
		return
	}
	defer db.Close()
	repo := repository.NewBooksRepo(db)
	books, err := repo.FetchByUser(userID)
	if err != nil {
		responses.Err(w, 500, err)
		return
	}
	responses.JSON(
		w,
		http.StatusOK,
		books,
	)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID, _ := strconv.Atoi(
		params["bookID"],
	)
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		responses.Err(w, 400, err)
		return
	}
	book.ID = bookID
	db, err := persistency.Connect()
	if err != nil {
		responses.Err(w, 500, err)
		return
	}
	defer db.Close()
	repo := repository.NewBooksRepo(db)
	err = repo.Update(book)
	if err != nil {
		responses.Err(w, 500, err)
		return
	}
	responses.JSON(
		w,
		http.StatusOK,
		map[string]string{
			"message": "livro atualizado",
		},
	)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID, _ := strconv.Atoi(
		params["bookID"],
	)
	db, err := persistency.Connect()
	if err != nil {
		responses.Err(w, 500, err)
		return
	}
	defer db.Close()
	repo := repository.NewBooksRepo(db)
	err = repo.Delete(bookID)
	if err != nil {
		responses.Err(w, 500, err)
		return
	}
	responses.JSON(
		w,
		http.StatusOK,
		map[string]string{
			"message": "livro removido",
		},
	)
}
