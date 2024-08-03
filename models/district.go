package models

type District struct {
	DistrictID   int64  `json:"district_id"`
	DistrictName string `json:"district_name"`
	ProvinceName string `json:"province_name"`
}

func (District) TableName() string {
	return "District"
}
