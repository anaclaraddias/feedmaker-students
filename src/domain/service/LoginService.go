package service

import (
	"errors"

	"github.com/anaclaraddias/feedmaker-students/src/domain/entity"
	"gorm.io/gorm"
)

const invalidLoginError = "incorrect username or password"

type LoginService struct {
	db *gorm.DB
}

func NewLoginService(db *gorm.DB) *LoginService {
	return &LoginService{
		db: db,
	}
}

func (service *LoginService) Execute(loginUser entity.User) (*entity.User, error) {
	var user *entity.User
	if err := service.db.Where("username = ?", loginUser.Username).Find(&user).Error; err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New(invalidLoginError)
	}

	if user.Password != loginUser.Password {
		return nil, errors.New(invalidLoginError)
	}

	return user, nil
}
