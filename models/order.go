package atest

type Order struct {
	Id      int    `json:"-"`
	Service int    `json:"serviceId"`
	User    int    `json:"userId"`
	Date    string `json:"date"`
	Status  bool   `json:"status"`
}
