package models

import (
	"gorm.io/gorm"
	"time"
)

type Employee struct {
	EmployeeID             int64          `json:"employee_id" gorm:"primaryKey;not null"`
	EmployeeName           string         `json:"employee_name" gorm:"not null"`
	EmployeeSurname        string         `json:"employee_surname" gorm:"not null"`
	EmployeeIdentityNumber int            `json:"employee_identity_number" gorm:"not null"`
	EmployeePhoneNumber    string         `json:"employee_phone_number" gorm:"not null"`
	EmployeeJobGroupName   string         `json:"employee_job_group_name" gorm:"not null"`
	EmployeeTitleName      string         `json:"employee_title_name" gorm:"not null"`
	EmployeeClinicName     string         `json:"employee_clinic_name" gorm:"not null"`
	EmployeeHospitalID     int64          `json:"employee_hospital_id" gorm:"not null"`
	CreatedAt              time.Time      `json:"created_at" gorm:"not null;default:now()"`
	UpdatedAt              time.Time      `json:"updated_at"`
	DeletedAt              gorm.DeletedAt `json:"deleted_at"`

	EmployeeWorkDays []EmployeeWorkDay `json:"EmployeeWorkDays" gorm:"foreignKey:EmployeeID;references:EmployeeID"`
}

func (Employee) TableName() string {
	return "Employee"
}
