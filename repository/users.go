package repository

import (
	"database/sql"
	"extensao-api/models"
)

type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{db}
}

func (u UsersRepo) Create(user models.Users) (int8, error) {
	query := `INSERT INTO treehousedb.users(
                            name,
                            email,
                            password,
                            cpf)
                            VALUES (?,?,?,?)`

	statement, err := u.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Email, user.Password, user.CPF)
	if err != nil {
		return 0, err
	}
	lastid, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int8(uint64(lastid)), nil
}

func (u UsersRepo) FetchByEmail(email string) (models.Users, error) {

	var user models.Users

	query := `
	SELECT
		id,
		email,
		password
	FROM users
	WHERE email = ?
	`

	err := u.db.QueryRow(
		query,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	)

	return user, err
}
