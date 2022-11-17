package repository

import (
	"fmt"

	balance "github.com/RomanMRR/user-balance"
	"github.com/jmoiron/sqlx"
)


type BalanceWalletPostgres struct {
	db *sqlx.DB
}

func NewBalanceWalletPostgres(db *sqlx.DB) *BalanceWalletPostgres {
	return &BalanceWalletPostgres{db:db}
}

// func (r *BalanceWalletPostgres) Update(userId int)

func (r *BalanceWalletPostgres) GetWallet(userId int) (balance.Wallet, error) {
	var wallet balance.Wallet
	query := fmt.Sprintf("SELECT wl.id, wl.amount, wl.user_id FROM %s wl INNER JOIN %s us on  wl.user_id=us.id WHERE us.id = $1", 
		walletTable, userTable)	
	err := r.db.Get(&wallet, query, userId)

	return wallet, err
}