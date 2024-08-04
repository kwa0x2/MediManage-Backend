package services

import (
	"github.com/kwa0x2/MediManage-Backend/models"
	"github.com/kwa0x2/MediManage-Backend/repositories"
	"gorm.io/gorm"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func (s *UserService) GetAll() ([]*models.User, error) {
	return s.UserRepository.GetAll()
}

func (s *UserService) Create(tx *gorm.DB, user *models.User) error {
	return s.UserRepository.Create(tx, user)
}

func (s *UserService) Delete(userId string) error {
	return s.UserRepository.Delete(userId)
}

func (s *UserService) Update(user *models.User, userId string) error {
	return s.UserRepository.Update(user, userId)
}

func (s *UserService) GetUserPasswordByIdentifier(identifier string) (*models.User, error) {
	return s.UserRepository.GetUserPasswordByIdentifier(identifier)
}
