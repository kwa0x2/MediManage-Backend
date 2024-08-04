package models

type Title struct {
	TitleID      int64  `json:"title_id" gorm:"primaryKey;not null"`
	TitleName    string `json:"title_name" gorm:"not null"`
	JobGroupName string `json:"jobgroup_name" gorm:"not null"`
}

func (Title) TableName() string {
	return "Title"
}
