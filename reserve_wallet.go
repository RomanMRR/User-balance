package balance

import "errors"

type ReserveWallet struct {
	Id int `json:"-" db:"id"`
	Amount float64 `json:"balance" db:"amount"`
	UserId int `json:"user_id" db:"user_id"`

}

type UpdateReserveWallet struct {
	User_id *int `json:"user_id"`
	Amount *float64 `json:"amount"`
	Service_id       *int  `json:"service_id" binding:"required"`
	OrderId       *int  `json:"order_id" binding:"required"`
}

func (i UpdateReserveWallet) Validate() error {
	if i.Amount == nil || i.User_id == nil {
		return errors.New("update structure has no values")
	}

	return nil
}