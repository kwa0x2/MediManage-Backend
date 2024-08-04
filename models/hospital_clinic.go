package models

import (
	"gorm.io/gorm"
	"time"
)

type HospitalClinic struct {
	HospitalID int64          `json:"hospital_id" gorm:"primaryKey;not null"`
	ClinicName string         `json:"clinic_name" gorm:"not null;primaryKey;"`
	CreatedAt  time.Time      `json:"created_at" gorm:"not null;default:now()"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

func (HospitalClinic) TableName() string {
	return "HospitalClinic"
}
