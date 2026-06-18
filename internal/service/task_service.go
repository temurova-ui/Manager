package service

import (
	"manager/internal/model"
	"manager/internal/repository"
)

type ITaskService interface {
	Create(task *model.Task) error
	GetAll() ([]model.Task, error)
	GetByID(id uint) (*model.Task, error)
	Update(task *model.Task) error
	Delete(id uint) error
}

type service struct {
	repo repository.ITaskRepo
}

func New(repo repository.ITaskRepo) ITaskService {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(task *model.Task) error {
	return s.repo.Create(task)
}

func (s *service) GetAll() ([]model.Task, error) {
	return s.repo.GetAll()
}

func (s *service) GetByID(id uint) (*model.Task, error) {
	return s.repo.GetByID(id)
}

func (s *service) Update(task *model.Task) error {
	return s.repo.Update(task)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}