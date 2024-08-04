package services

import (
	"github.com/kwa0x2/MediManage-Backend/models"
	"github.com/kwa0x2/MediManage-Backend/repositories"
	"github.com/kwa0x2/MediManage-Backend/types"
	"gorm.io/gorm"
)

type EmployeeService struct {
	EmployeeRepository *repositories.EmployeeRepository
}

func (s *EmployeeService) GetByHospitalId(hospitalId int64) ([]*models.Employee, error) {
	return s.EmployeeRepository.GetByHospitalId(hospitalId)
}

func (s *EmployeeService) Delete(employeeId int64) error {
	tx := s.EmployeeRepository.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := s.EmployeeRepository.Delete(tx, employeeId); err != nil {
		tx.Rollback()
		return err
	}

	if err := s.EmployeeRepository.DeleteWorkDays(tx, employeeId); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil

}

func (s *EmployeeService) Update(employee *models.Employee, employeeWorkingDays []string, employeeId int64) error {
	tx := s.EmployeeRepository.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := s.EmployeeRepository.Update(tx, employee, employeeId); err != nil {
		tx.Rollback()
		return err
	}

	if err := s.EmployeeRepository.DeleteWorkDays(tx, employeeId); err != nil {
		tx.Rollback()
		return err
	}

	for _, dayStr := range employeeWorkingDays {
		employeeWorkDay := models.EmployeeWorkDay{
			EmployeeID: employeeId,
			Day:        types.DayType(dayStr),
		}

		if err := s.EmployeeRepository.CreateWorkDay(tx, &employeeWorkDay); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (s *EmployeeService) CheckChiefDoctorExists(hospitalId int64) (bool, error) {
	return s.EmployeeRepository.IsChiefDoctorExists(hospitalId)
}

func (s *EmployeeService) Create(employee *models.Employee, employeeWorkingDays []string) error {
	tx := s.EmployeeRepository.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	addedEmployeeData, err := s.EmployeeRepository.Create(tx, employee)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, dayStr := range employeeWorkingDays {
		employeeWorkDay := models.EmployeeWorkDay{
			EmployeeID: addedEmployeeData.EmployeeID,
			Day:        types.DayType(dayStr),
		}

		if err := s.EmployeeRepository.CreateWorkDay(tx, &employeeWorkDay); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (s *EmployeeService) DeleteByClinicName(tx *gorm.DB, clinicName int64) error {
	if err := s.EmployeeRepository.Delete(tx, clinicName); err != nil {
		return err
	}

	return nil
}
