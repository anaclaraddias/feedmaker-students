package service

import (
	"github.com/anaclaraddias/feedmaker-students/src/domain/entity"
	"gorm.io/gorm"
)

type CreateUserService struct {
	db *gorm.DB
}

func NewCreateUserService(db *gorm.DB) *CreateUserService {
	return &CreateUserService{
		db: db,
	}
}

func (service *CreateUserService) Execute(student *entity.User) error {
	if err := service.db.Create(student).Error; err != nil {
		return err
	}

	return nil
}
