package services

import (
	"github.com/kwa0x2/MediManage-Backend/models"
	"github.com/kwa0x2/MediManage-Backend/repositories"
)

type DistrictService struct {
	DistrictRepository *repositories.DistrictRepository
}

func (s *DistrictService) GetAll() ([]*models.District, error) {
	return s.DistrictRepository.GetAll()
}

func (s *DistrictService) GetAllByProvinceName(provinceName string) ([]*models.District, error) {
	return s.DistrictRepository.GetAllByProvinceName(provinceName)
}
