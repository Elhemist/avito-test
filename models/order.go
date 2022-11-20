package atest

type Order struct {
	Id      int    `json:"orderId"`
	Service int    `json:"serviceId" binding:"required"`
	User    int    `json:"userId" binding:"required"`
	Date    string `json:"date"`
	Status  bool   `json:"status"`
}
type OrderResp struct {
	Id      int `json:"orderId"`
	Service int `json:"serviceId" binding:"required"`
	User    int `json:"userId" binding:"required"`
	Prise   int `json:"prise"`
}
