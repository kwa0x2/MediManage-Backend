package models

import (
	"github.com/google/uuid"
	"github.com/kwa0x2/MediManage-Backend/types"
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserID             uuid.UUID          `json:"user_id" gorm:"primaryKey;not null;type:uuid;default:gen_random_uuid()"`
	UserName           string             `json:"user_name" gorm:"not null"`
	UserSurname        string             `json:"user_surname" gorm:"not null"`
	UserIdentityNumber string             `json:"user_identity_number" gorm:"not null;size:11"`
	UserEmail          string             `json:"user_email" gorm:"not null"`
	UserPhone          string             `json:"user_phone" gorm:"not null;size:15"`
	UserPassword       string             `json:"user_password" gorm:"not null"`
	UserRole           types.UserRoleType `json:"user_role" gorm:"not null"`
	CreatedAt          time.Time          `json:"created_at" gorm:"not null;default:now()"`
	UpdatedAt          time.Time          `json:"updated_at"`
	DeletedAt          gorm.DeletedAt     `json:"deleted_at"`
}

func (User) TableName() string {
	return "User"
}
