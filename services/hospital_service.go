package services

import (
	"github.com/kwa0x2/MediManage-Backend/models"
	"github.com/kwa0x2/MediManage-Backend/repositories"
	"gorm.io/gorm"
)

type HospitalService struct {
	HospitalRepository *repositories.HospitalRepository
}

func (s *HospitalService) Create(tx *gorm.DB, hospital *models.Hospital) error {
	return s.HospitalRepository.Create(tx, hospital)
}
