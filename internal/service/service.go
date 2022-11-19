package service

import (
	repository "github.com/Elhemist/avito-test/internal/repositiry"
	atest "github.com/Elhemist/avito-test/models"
)

type Order interface {
	CreateOrder(order atest.Order) (int, error)
	RevenueOrder(order atest.Order) (int, error)
}
type User interface {
	CheckUser(user atest.User) (atest.User, error)
	UpdateUser(user atest.User) (atest.User, error)
}

type Service struct {
	Order
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Order: NewOrderService(repos.Order),
		User:  NewUserService(repos.User),
	}
}
