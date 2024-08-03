package models

import (
	"gorm.io/gorm"
	"time"
)

type Hospital struct {
	HospitalID                int64          `json:"hospital_id" gorm:"primaryKey;not null"`
	HospitalName              string         `json:"hospital_name" gorm:"not null"`
	HospitalTaxIdentityNumber string         `json:"hospital_tax_identity_number" gorm:"not null;size:10"`
	HospitalEmail             string         `json:"hospital_email" gorm:"not null"`
	HospitalPhoneNumber       string         `json:"hospital_phone_number" gorm:"not null;size:15"`
	HospitalProvinceName      string         `json:"hospital_province_name" gorm:"not null"`
	HospitalDistrictName      string         `json:"hospital_district_name" gorm:"not null"`
	HospitalAddress           string         `json:"hospital_address" gorm:"not null"`
	CreatedAt                 time.Time      `json:"created_at" gorm:"not null;default:now()"`
	UpdatedAt                 time.Time      `json:"updated_at"`
	DeletedAt                 gorm.DeletedAt `json:"deleted_at"`
}

func (Hospital) TableName() string {
	return "Hospital"
}
