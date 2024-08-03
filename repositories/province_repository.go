package repositories

import (
	"github.com/kwa0x2/MediManage-Backend/models"
	"gorm.io/gorm"
)

type ProvinceRepository struct {
	DB *gorm.DB
}

func (r *ProvinceRepository) GetAll() ([]*models.Province, error) {
	var provinces []*models.Province
	if err := r.DB.Find(&provinces).Error; err != nil {
		return nil, err
	}

	return provinces, nil
}
