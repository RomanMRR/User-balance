package balance

type Wallet struct {
	Id int `json:"-" db:"id"`
	Amount float64 `json:"balance" db:"amount"`
	UserId int `json:"user_id" db:"user_id"`
}
