package repositories

import (
	"github.com/kwa0x2/MediManage-Backend/models"
	"gorm.io/gorm"
)

type HospitalRepository struct {
	DB *gorm.DB
}

func (r *HospitalRepository) Create(tx *gorm.DB, hospital *models.Hospital) error {
	db := r.DB
	if tx != nil {
		db = tx
	}

	if err := db.Create(&hospital).Error; err != nil {
		return err
	}
	return nil
}
