package repositories

import (
	"github.com/kwa0x2/MediManage-Backend/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Create(tx *gorm.DB, user *models.User) error {
	db := r.DB
	if tx != nil {
		db = tx
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserPasswordByIdentifier(identifier string) (*models.User, error) {
	var user *models.User
	err := r.DB.Model(&models.User{}).Debug().
		Where("user_email = ? OR user_phone = ?", identifier, identifier).
		First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
