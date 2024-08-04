package services

import (
	"github.com/kwa0x2/MediManage-Backend/models"
	"github.com/kwa0x2/MediManage-Backend/repositories"
)

type TitleService struct {
	TitleRepository *repositories.TitleRepository
}

func (s *TitleService) GetAllByJobGroupName(jobGroupName string) ([]*models.Title, error) {
	return s.TitleRepository.GetAllByJobGroupName(jobGroupName)
}
