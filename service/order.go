package service

import (
	"github.com/train-do/Golang-Generics/model"
	"github.com/train-do/Golang-Generics/repository"
)

type ServiceOrder struct {
	Repo *repository.RepoOrder
}

func NewServiceOrder(repo *repository.RepoOrder) *ServiceOrder {
	return &ServiceOrder{repo}
}
func (s *ServiceOrder) AddOrder(order *model.Order) error {
	err := s.Repo.Insert(order)
	if err != nil {
		return err
	}
	return nil
}
