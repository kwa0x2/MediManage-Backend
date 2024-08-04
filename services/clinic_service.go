package services

import (
	"github.com/kwa0x2/MediManage-Backend/models"
	"github.com/kwa0x2/MediManage-Backend/repositories"
)

type ClinicService struct {
	ClinicRepository   *repositories.ClinicRepository
	EmployeeRepository *repositories.EmployeeRepository
}

func (s *ClinicService) GetAll() ([]*models.Clinic, error) {
	return s.ClinicRepository.GetAll()
}

func (s *ClinicService) GetAllHospitalClinic(hospitalId int64) ([]map[string]interface{}, error) {
	return s.ClinicRepository.GetAllHospitalClinic(hospitalId)
}

func (s *ClinicService) CreateHospitalClinic(hospitalClinic []string, hospitalId int64) error {
	tx := s.ClinicRepository.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	for _, hospitalClinicStr := range hospitalClinic {
		hospitalClinicObj := models.HospitalClinic{
			HospitalID: hospitalId,
			ClinicName: hospitalClinicStr,
		}

		if err := s.ClinicRepository.CreateHospitalClinic(tx, &hospitalClinicObj); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (s *ClinicService) DeleteHospitalClinicByClinicName(hospitalId int64, clinicName string) error {
	tx := s.ClinicRepository.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := s.ClinicRepository.DeleteHospitalClinicByClinicName(tx, hospitalId, clinicName); err != nil {
		tx.Rollback()
		return err
	}

	if err := s.EmployeeRepository.DeleteByClinicName(tx, clinicName); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (s *ClinicService) Update(hospitalClinic []string, hospitalId int64) error {
	tx := s.ClinicRepository.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := s.ClinicRepository.DeleteHospitalClinic(tx, hospitalId); err != nil {
		tx.Rollback()
		return err
	}

	if err := s.ClinicRepository.DeleteHospitalClinic(tx, hospitalId); err != nil {
		tx.Rollback()
		return err
	}

	for _, hospitalClinicStr := range hospitalClinic {
		hospitalClinicObj := models.HospitalClinic{
			HospitalID: hospitalId,
			ClinicName: hospitalClinicStr,
		}

		if err := s.ClinicRepository.CreateHospitalClinic(tx, &hospitalClinicObj); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
