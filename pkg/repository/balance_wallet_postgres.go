package repository

import (
	"fmt"
	"strings"

	balance "github.com/RomanMRR/user-balance"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)


type BalanceWalletPostgres struct {
	db *sqlx.DB
}

func NewBalanceWalletPostgres(db *sqlx.DB) *BalanceWalletPostgres {
	return &BalanceWalletPostgres{db:db}
}

func (r *BalanceWalletPostgres) Update(input balance.UpadateWallet) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	var count int
	queryCount := fmt.Sprintf("SELECT count(id) FROM %s WHERE user_id = $1", walletTable)
	row := r.db.QueryRow(queryCount, input.User_id)
	if err := row.Scan(&count); err != nil {
		return err
	}

	if count != 0 {
		if input.Amount != nil {
			setValues = append(setValues, fmt.Sprintf("amount=amount + $%v", argId))
			args = append(args, *input.Amount)
			argId++
		}
	
		setQuery := strings.Join(setValues, ", ")
	
		query := fmt.Sprintf("UPDATE %s wl SET %s FROM %s us WHERE wl.user_id = us.id AND us.id=$%d",
			walletTable, setQuery, userTable, argId)	
		args = append(args, input.User_id)
		
		logrus.Debugf("updateQuery: %s", query)
		logrus.Debugf("args: %s", args)
	
		_, err := r.db.Exec(query, args...)
		return err
	} else {
		tx, err := r.db.Begin()
		if err != nil {
			return err
		}

		createWalletQuery := fmt.Sprintf("INSERT INTO %s (user_id, amount) VALUES($1, $2)", walletTable)
		_, err = tx.Exec(createWalletQuery, input.User_id, input.Amount)
		if err != nil {
			tx.Rollback()
			return err
		}

		return tx.Commit()
		}	
}

func (r *BalanceWalletPostgres) GetWallet(userId int) (balance.Wallet, error) {
	var wallet balance.Wallet
	query := fmt.Sprintf("SELECT wl.id, wl.amount, wl.user_id FROM %s wl INNER JOIN %s us on  wl.user_id=us.id WHERE us.id = $1", 
		walletTable, userTable)	
	err := r.db.Get(&wallet, query, userId)

	return wallet, err
}