package balance

type User struct {
	Id             int     `json:"-" db:"id"`
	Name           string  `json:"name" binding:"required"`
	Username       string  `json:"username" binding:"required"`
	Password       string  `json:"password" binding:"required"`
	Balance        float64 `json:"balance"`
	ReserveBalance float64 `json:"reserve-balance"`
}
