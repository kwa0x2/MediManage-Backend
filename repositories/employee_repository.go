package repositories

import (
	"github.com/kwa0x2/MediManage-Backend/models"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

func (r *EmployeeRepository) GetByHospitalId(hospitalId int64) ([]*models.Employee, error) {
	var employees []*models.Employee
	if err := r.DB.Where("employee_hospital_id = ?", hospitalId).Preload("EmployeeWorkDays").Find(&employees).Error; err != nil {
		return nil, err
	}

	return employees, nil
}

func (r *EmployeeRepository) Create(tx *gorm.DB, employee *models.Employee) (*models.Employee, error) {
	db := r.DB
	if tx != nil {
		db = tx
	}

	if err := db.Create(&employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
}

func (r *EmployeeRepository) IsChiefDoctorExists(hospitalID int64) (bool, error) {
	var count int64
	err := r.DB.Model(&models.Employee{}).
		Where("employee_title_name = ? AND employee_hospital_id = ?", "Başhekim", hospitalID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *EmployeeRepository) CreateWorkDay(tx *gorm.DB, employeeWorkDay *models.EmployeeWorkDay) error {
	db := r.DB
	if tx != nil {
		db = tx
	}

	if err := db.Create(&employeeWorkDay).Error; err != nil {
		return err
	}
	return nil
}

func (r *EmployeeRepository) Delete(tx *gorm.DB, employeeId int64) error {
	db := r.DB
	if tx != nil {
		db = tx
	}

	if err := db.Delete(&models.Employee{}, "employee_id = ?", employeeId).Error; err != nil {
		return err
	}
	return nil
}

func (r *EmployeeRepository) DeleteWorkDays(tx *gorm.DB, employeeId int64) error {
	db := r.DB
	if tx != nil {
		db = tx
	}

	if err := db.Delete(&models.EmployeeWorkDay{}, "employee_id = ?", employeeId).Error; err != nil {
		return err
	}
	return nil
}

func (r *EmployeeRepository) Update(tx *gorm.DB, employee *models.Employee, employeeId int64) error {
	db := r.DB
	if tx != nil {
		db = tx
	}

	if err := db.Model(&employee).Where("employee_id = ?", employeeId).Updates(employee).Error; err != nil {
		return err
	}
	return nil
}

func (r *EmployeeRepository) DeleteByClinicName(tx *gorm.DB, clinicName string) error {
	db := r.DB
	if tx != nil {
		db = tx
	}

	if err := db.Delete(&models.Employee{}, "employee_clinic_name = ?", clinicName).Error; err != nil {
		return err
	}
	return nil
}
