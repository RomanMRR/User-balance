package repository

import (
	"fmt"

	balance "github.com/RomanMRR/user-balance"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user balance.User) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createUserQuery := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", userTable)
	row := tx.QueryRow(createUserQuery, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createReserveWalletQuery := fmt.Sprintf("INSERT INTO %s (user_id) VALUES($1)", reserveWalletTable)
	_, err = tx.Exec(createReserveWalletQuery, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *AuthPostgres) GetUser(username, password string) (balance.User, error) {
	var user balance.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", userTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}