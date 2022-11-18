package atest

type Service struct {
	Id    int    `json:"-"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
