package repository

import (
	"manager/internal/model"

	"gorm.io/gorm"
)

type ITaskRepo interface {
	Create(task *model.Task) error
	GetAll() ([]model.Task, error)
	GetByID(id int) (*model.Task, error)
	Update(task *model.Task) error
	Delete(id int) error
}

type repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ITaskRepo {
	return &repo{db: db}
}

func (r *repo) Create(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *repo) GetAll() ([]model.Task, error) {
	var tasks []model.Task

	err := r.db.Find(&tasks).Error

	return tasks, err
}

func (r *repo) GetByID(id int) (*model.Task, error) {
	var task model.Task

	err := r.db.First(&task, id).Error

	return &task, err
}

func (r *repo) Update(task *model.Task) error {
	return r.db.Save(task).Error
}

func (r *repo) Delete(id int) error {
	return r.db.Delete(&model.Task{}, id).Error
}