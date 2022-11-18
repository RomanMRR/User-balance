package balance

import "errors"

type Wallet struct {
	Id int `json:"-" db:"id"`
	Amount float64 `json:"balance" db:"amount"`
	UserId int `json:"user_id" db:"user_id"`
}

type UpdateWallet struct {
	User_id *int `json:"user_id"`
	Amount *float64 `json:"amount"`
}

func (i UpdateWallet) Validate() error {
	if i.Amount == nil || i.User_id == nil {
		return errors.New("update structure has no values")
	}

	return nil
}