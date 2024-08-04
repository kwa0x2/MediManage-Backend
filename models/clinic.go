package models

type Clinic struct {
	ClinicID   int64  `json:"clinic_id" gorm:"primaryKey;not null"`
	ClinicName string `json:"clinic_name" gorm:"not null"`
}

func (Clinic) TableName() string {
	return "Clinic"
}
