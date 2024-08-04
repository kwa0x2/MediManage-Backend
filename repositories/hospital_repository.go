package repositories

import (
	"github.com/kwa0x2/MediManage-Backend/models"
	"gorm.io/gorm"
)

type HospitalRepository struct {
	DB *gorm.DB
}

func (r *HospitalRepository) Create(tx *gorm.DB, hospital *models.Hospital) (int64, error) {
	db := r.DB
	if tx != nil {
		db = tx
	}

	if err := db.Create(&hospital).Error; err != nil {
		return 0, err
	}
	return hospital.HospitalID, nil
}

func (r *HospitalRepository) GetById(hospitalId int64) (*models.Hospital, error) {
	var hospital *models.Hospital
	err := r.DB.Model(&models.Hospital{}).
		Where("hospital_id = ? ", hospitalId).
		First(&hospital).Error
	if err != nil {
		return nil, err
	}
	return hospital, nil
}
