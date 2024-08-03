package repositories

import (
	"github.com/kwa0x2/MediManage-Backend/models"
	"gorm.io/gorm"
)

type DistrictRepository struct {
	DB *gorm.DB
}

func (r *DistrictRepository) GetAll() ([]*models.District, error) {
	var districts []*models.District
	if err := r.DB.Find(&districts).Error; err != nil {
		return nil, err
	}

	return districts, nil
}

func (r *DistrictRepository) GetAllByProvinceName(provinceName string) ([]*models.District, error) {
	var districts []*models.District
	if err := r.DB.Where("province_name", provinceName).Find(&districts).Error; err != nil {
		return nil, err
	}

	return districts, nil
}
