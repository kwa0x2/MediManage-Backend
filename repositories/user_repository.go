package repositories

import (
	"github.com/kwa0x2/MediManage-Backend/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) GetAll() ([]*models.User, error) {
	var users []*models.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
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

func (r *UserRepository) Delete(userId string) error {
	if err := r.DB.Delete(&models.User{}, "user_id = ?", userId).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Update(user *models.User, userId string) error {
	if err := r.DB.Model(&user).Where("user_id = ?", userId).Updates(user).Error; err != nil {
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
