package balance

import "time"

type Transaction struct {
	Id        int       `json:"id"`
	ServiceId int       `json:"service-id"`
	OrderId   int       `json:"order-id"`
	Sum       float64   `json:"Sum"`
	Date      time.Time `json:"Date"`
}
