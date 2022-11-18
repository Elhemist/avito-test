package atest

type Service struct {
	Id    int    `json:"serviceId"`
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required"`
}
