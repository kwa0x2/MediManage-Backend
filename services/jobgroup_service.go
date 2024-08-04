package services

import (
	"github.com/kwa0x2/MediManage-Backend/models"
	"github.com/kwa0x2/MediManage-Backend/repositories"
)

type JobGroupService struct {
	JobGroupRepository *repositories.JobGroupRepository
}

func (s *JobGroupService) GetAll() ([]*models.JobGroup, error) {
	return s.JobGroupRepository.GetAll()
}
