package models

type JobGroup struct {
	JobGroupID   int64  `json:"jobgroup_id" gorm:"primaryKey;not null:jobgroup_id"`
	JobGroupName string `json:"jobgroup_name" gorm:"not null;column:jobgroup_name"`
}

func (JobGroup) TableName() string {
	return "JobGroup"
}
