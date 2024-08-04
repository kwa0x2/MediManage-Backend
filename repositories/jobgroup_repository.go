package repositories

import (
	"github.com/kwa0x2/MediManage-Backend/models"
	"gorm.io/gorm"
)

type JobGroupRepository struct {
	DB *gorm.DB
}

func (r *JobGroupRepository) GetAll() ([]*models.JobGroup, error) {
	var jobGroups []*models.JobGroup
	if err := r.DB.Debug().Find(&jobGroups).Error; err != nil {
		return nil, err
	}

	return jobGroups, nil
}
