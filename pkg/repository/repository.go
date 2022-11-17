package repository

import (
	balance "github.com/RomanMRR/user-balance"
	"github.com/jmoiron/sqlx"
)


type Authorization interface {
	CreateUser(user balance.User)(int, error)
	GetUser(username, password string) (balance.User, error)
}

type BalanceWallet interface {
	// Update(userId int, wallet balance.Wallet)
	GetWallet(userId int)(balance.Wallet, error)
}

type Repository struct {
	Authorization
	BalanceWallet
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		BalanceWallet: NewBalanceWalletPostgres(db),
	}
}