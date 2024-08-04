package repositories

import (
	"github.com/kwa0x2/MediManage-Backend/models"
	"gorm.io/gorm"
)

type ClinicRepository struct {
	DB *gorm.DB
}

func (r *ClinicRepository) GetAll() ([]*models.Clinic, error) {
	var clinics []*models.Clinic
	if err := r.DB.Find(&clinics).Error; err != nil {
		return nil, err
	}

	return clinics, nil
}

func (r *ClinicRepository) CreateHospitalClinic(tx *gorm.DB, hospitalClinic *models.HospitalClinic) error {
	db := r.DB
	if tx != nil {
		db = tx
	}

	if err := db.Model(&models.HospitalClinic{}).Create(&hospitalClinic).Error; err != nil {
		return err
	}
	return nil
}

func (r *ClinicRepository) GetAllHospitalClinic(hospitalId int64) ([]map[string]interface{}, error) {
	var hospitalClinics []map[string]interface{}

	if err := r.DB.Model(&models.HospitalClinic{}).
		Select(`
            "HospitalClinic".clinic_name,
            COUNT("Employee".employee_id) AS total_employees,
            COALESCE(SUM(CASE WHEN "Employee".employee_job_group_name = 'Doktor' THEN 1 ELSE 0 END), 0) AS doktor_count,
            COALESCE(SUM(CASE WHEN "Employee".employee_job_group_name = 'Hizmet Personeli' THEN 1 ELSE 0 END), 0) AS hizmet_personeli_count,
            COALESCE(SUM(CASE WHEN "Employee".employee_job_group_name = 'İdari Personel' THEN 1 ELSE 0 END), 0) AS idari_personel_count
        `).
		Joins("LEFT JOIN \"Employee\" ON \"HospitalClinic\".clinic_name = \"Employee\".employee_clinic_name AND \"HospitalClinic\".hospital_id = \"Employee\".employee_hospital_id AND \"Employee\".deleted_at IS NULL").
		Where("\"HospitalClinic\".hospital_id = ?", hospitalId).
		Group("\"HospitalClinic\".clinic_name").
		Order("\"HospitalClinic\".clinic_name").
		Scan(&hospitalClinics).Error; err != nil {
		return nil, err
	}

	return hospitalClinics, nil
}

func (r *ClinicRepository) DeleteHospitalClinicByClinicName(tx *gorm.DB, hospitalId int64, clinicName string) error {
	db := r.DB
	if tx != nil {
		db = tx
	}

	if err := db.Delete(&models.HospitalClinic{}, "hospital_id = ? AND clinic_name = ? ", hospitalId, clinicName).Error; err != nil {
		return err
	}
	return nil
}

func (r *ClinicRepository) DeleteHospitalClinic(tx *gorm.DB, hospitalId int64) error {
	db := r.DB
	if tx != nil {
		db = tx
	}

	if err := db.Delete(&models.HospitalClinic{}, "hospital_id = ?", hospitalId).Error; err != nil {
		return err
	}
	return nil
}
