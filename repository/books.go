package repository

import (
	"database/sql"
	"extensao-api/models"
)

type BooksRepo struct {
	db *sql.DB
}

func NewBooksRepo(db *sql.DB) *BooksRepo {
	return &BooksRepo{db}
}

func (b BooksRepo) Create(book models.Book) error {

	query := `
	INSERT INTO books
	(
	google_book_id,
	title,
	authors,
	description,
	user_id
	)
	VALUES (?,?,?,?,?)
	`
	_, err := b.db.Exec(
		query,
		book.GoogleBookID,
		book.Title,
		book.Authors,
		book.Description,
		book.UserID,
	)
	return err
}

func (b BooksRepo) FetchByUser(userID int) ([]models.Book, error) {
	rows, err := b.db.Query(
		"SELECT * FROM books WHERE user_id=?",
		userID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []models.Book
	for rows.Next() {
		var book models.Book
		err = rows.Scan(
			&book.ID,
			&book.GoogleBookID,
			&book.Title,
			&book.Authors,
			&book.Description,
			&book.UserID,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (b BooksRepo) Delete(id int) error {
	_, err := b.db.Exec(
		"DELETE FROM books WHERE id=?",
		id,
	)
	return err
}

func (b BooksRepo) Update(book models.Book) error {
	query := `
	UPDATE books
	SET
		title=?,
		authors=?,
		description=?
	WHERE id=?
	`
	_, err := b.db.Exec(
		query,
		book.Title,
		book.Authors,
		book.Description,
		book.ID,
	)
	return err
}
