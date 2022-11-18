package atest

type User struct {
	Id         int `json:"userId"`
	Balance    int `json:"balance" binding:"required"`
	BlanceTemp int `json:"balanceTemp" binding:"required"`
}
