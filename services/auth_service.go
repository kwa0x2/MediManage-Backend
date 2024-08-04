package services

import (
	"github.com/kwa0x2/MediManage-Backend/models"
	"github.com/kwa0x2/MediManage-Backend/repositories"
)

type AuthService struct {
	UserRepository     *repositories.UserRepository
	HospitalRepository *repositories.HospitalRepository
}

func (s *AuthService) Register(hospital *models.Hospital, user *models.User) error {
	tx := s.HospitalRepository.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	addedHospitalId, err := s.HospitalRepository.Create(tx, hospital)
	if err != nil {
		tx.Rollback()
		return err
	}

	user.UserHospitalID = addedHospitalId

	if err := s.UserRepository.Create(tx, user); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
