package models

import (
	"github.com/kwa0x2/MediManage-Backend/types"
	"time"
)

type EmployeeWorkDay struct {
	EmployeeID int64         `json:"employee_id" gorm:"primaryKey;not null"`
	Day        types.DayType `json:"day" gorm:"not null;primaryKey;type:day_type"`
	CreatedAt  time.Time     `json:"created_at" gorm:"not null;default:now()"`
}

func (EmployeeWorkDay) TableName() string {
	return "EmployeeWorkDay"
}
