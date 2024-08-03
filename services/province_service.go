package services

import (
	"github.com/kwa0x2/MediManage-Backend/models"
	"github.com/kwa0x2/MediManage-Backend/repositories"
)

type ProvinceService struct {
	ProvinceRepository *repositories.ProvinceRepository
}

func (s *ProvinceService) GetAll() ([]*models.Province, error) {
	return s.ProvinceRepository.GetAll()
}
