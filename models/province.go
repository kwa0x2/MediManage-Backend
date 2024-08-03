package models

type Province struct {
	ProvinceID   int64  `json:"province_id"`
	ProvinceName string `json:"province_name"`
}

func (Province) TableName() string {
	return "Province"
}
