package service

import repository "github.com/Elhemist/avito-test/internal/repositiry"

type Order interface {
}
type User interface {
}

type Service struct {
	Order
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
