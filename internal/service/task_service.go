package service

import (
  "errors"
  "strings"
  "manager/internal/model"
  "manager/internal/repository"
)

type ItaskService interface {
  Create(task model.Task) error
  GetAll() ([]model.Task, error)
  GetByID(id int) (*model.Task, error)
  Delete(id int) error
}

type service struct {
  repo repository.ITaskRepo
}

func New(repo repository.ITaskRepo) ItaskService {
  return &service{
    repo: repo,
  }
}

func (s *service) Create(task model.Task) error {
  if strings.TrimSpace(task.Title) == "" {
    return errors.New("empty title")
  }
  return s.repo.Create(&task)
}

func (s *service) GetAll() ([]model.Task, error) {
  return s.repo.GetAll()
}

func (s *service) GetByID(id int) (*model.Task, error) {
  if id == 0 {
    return nil, errors.New("invalid id")
  }
  return s.repo.GetByID(id)
}

func (s *service) Delete(id int) error {
  if id == 0 {
    return errors.New("invalid id")
  }
  return s.repo.Delete(id)
}