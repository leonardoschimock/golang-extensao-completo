package models

type Book struct {
	ID           int    `json:"id"`
	GoogleBookID string `json:"google_book_id"`
	Title        string `json:"title"`
	Authors      string `json:"authors"`
	Description  string `json:"description"`
	UserID       int    `json:"user_id"`
}
