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
	Update(input balance.UpadateWallet) error
	GetWallet(userId int)(balance.Wallet, error)
}

type ReserveWallet interface {
	AddAmount(input balance.UpdateReserveWallet) error
	WithdrawAmount(input balance.UpdateReserveWallet) error
	GetReserveWallet(userId int) (balance.ReserveWallet, error)
}

type Repository struct {
	Authorization
	BalanceWallet
	ReserveWallet
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		BalanceWallet: NewBalanceWalletPostgres(db),
		ReserveWallet: NewReserveWalletPostgres(db),
	}
}