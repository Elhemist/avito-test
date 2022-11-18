package service

import (
	repository "github.com/Elhemist/avito-test/internal/repositiry"
	atest "github.com/Elhemist/avito-test/models"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CheckUser(User atest.User) (int, error) {
	return s.repo.CheckUser(User)
}
