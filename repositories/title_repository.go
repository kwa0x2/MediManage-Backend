package repositories

import (
	"github.com/kwa0x2/MediManage-Backend/models"
	"gorm.io/gorm"
)

type TitleRepository struct {
	DB *gorm.DB
}

func (r *TitleRepository) GetAllByJobGroupName(jobGroupName string) ([]*models.Title, error) {
	var titles []*models.Title
	if err := r.DB.Where("jobgroup_name = ?", jobGroupName).Find(&titles).Error; err != nil {
		return nil, err
	}

	return titles, nil
}
