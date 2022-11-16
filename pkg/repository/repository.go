package repository

import (
	balance "github.com/RomanMRR/user-balance"
	"github.com/jmoiron/sqlx"
)


type Authorization interface {
	CreateUser(user balance.User)(int, error)
	GetUser(username, password string) (balance.User, error)
}

type Balance interface {

}

type Repository struct {
	Authorization
	Balance
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}