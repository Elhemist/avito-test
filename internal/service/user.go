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

func (s *UserService) CheckUser(user atest.User) (atest.User, error) {

	return s.repo.CheckUser(user)
}
func (s *UserService) UpdateUser(user atest.User) (atest.User, error) {

	return s.repo.UpdateUser(user)
}
