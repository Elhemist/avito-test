package service

type Order interface {
}
type User interface {
}

type Service struct {
	Order
	User
}

func NewService() *Service {
	return &Service{}
}
