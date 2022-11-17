package repository

import (
	"fmt"

	balance "github.com/RomanMRR/user-balance"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)


type ReserveWalletPostgres struct {
	db *sqlx.DB
}

func NewReserveWalletPostgres(db *sqlx.DB) *ReserveWalletPostgres {
	return &ReserveWalletPostgres{db:db}
}

func (r *ReserveWalletPostgres) AddAmount(input balance.UpdateReserveWallet) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	addQuery := fmt.Sprintf("UPDATE %s rw SET amount=amount + $1 FROM %s us WHERE rw.user_id = us.id AND us.id=$2",
		reserveWalletTable, userTable)

	_, err = tx.Exec(addQuery, input.Amount, input.User_id)
	logrus.Debugf("addAmountQuery: %s", addQuery)
	if err != nil {
		tx.Rollback()
		return err
	}

	

	withdrawQuery := fmt.Sprintf("UPDATE %s wl SET amount=amount - $1 FROM %s us WHERE wl.user_id = us.id AND us.id=$2", 
			walletTable, userTable)
	_, err = tx.Exec(withdrawQuery, input.Amount, input.User_id)
	logrus.Debugf("withdrawQuery: %s", withdrawQuery)
	if err != nil {
		tx.Rollback()
		return err
	}
	

	return tx.Commit()
}

func (r *ReserveWalletPostgres) GetReserveWallet(userId int) (balance.ReserveWallet, error) {
	var reserveWallet balance.ReserveWallet
	query := fmt.Sprintf("SELECT rw.id, rw.amount, rw.user_id FROM %s rw INNER JOIN %s us on  rw.user_id=us.id WHERE us.id = $1", 
		reserveWalletTable, userTable)	
	err := r.db.Get(&reserveWallet, query, userId)

	return reserveWallet, err
}

func (r *ReserveWalletPostgres) WithdrawAmount(input balance.UpdateReserveWallet) error {
	queryWithdraw := fmt.Sprintf("UPDATE %s rw SET amount=amount - $1 FROM %s us WHERE rw.user_id = us.id AND us.id=$2",
	reserveWalletTable, userTable)
	
	_, err := r.db.Exec(queryWithdraw, input.Amount, input.User_id)
	logrus.Debugf("queryWithdraw: %s", queryWithdraw)
	return err

}