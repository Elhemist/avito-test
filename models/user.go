package atest

type User struct {
	Id         int `json:"-"`
	Balance    int `json:"balance"`
	BlanceTemp int `json:"balanceTemp"`
}
