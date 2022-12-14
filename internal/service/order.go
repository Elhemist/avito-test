package service

import (
	"fmt"

	repository "github.com/Elhemist/avito-test/internal/repositiry"
	atest "github.com/Elhemist/avito-test/models"
)

	
	type OrderService struct {
		repo repository.Order
	}
	
	func NewOrderService(repo repository.Order) *OrderService {
		return &OrderService{repo: repo}
	}
	func (s *OrderService) CreateOrder(order atest.Order) (int, error) {
		return s.repo.CreateOrder(order)
	}
	func (s *OrderService) RevenueOrder(order atest.Order) (int, error) {
		fmt.Println("nya2")
		return s.repo.RevenueOrder(order)
	}

	
