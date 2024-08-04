package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/MediManage-Backend/cache"
	"github.com/kwa0x2/MediManage-Backend/services"
	"net/http"
)

type ClinicController struct {
	ClinicService *services.ClinicService
	ClinicCache   *cache.ClinicCache
}

func (c *ClinicController) GetAll(ctx *gin.Context) {
	clinics, err := c.ClinicCache.GetAllClinics()
	if clinics == nil && len(clinics) > 0 {
		ctx.JSON(http.StatusOK, clinics)
		return
	}

	clinics, err = c.ClinicService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error fetching clinics")
		return
	}

	err = c.ClinicCache.SetAllClinics(clinics)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, clinics)
}

func (c *ClinicController) GetAllHospitalClinic(ctx *gin.Context) {
	session := sessions.Default(ctx)

	userHospitalID := session.Get("hospital_id")
	if userHospitalID == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "UserHospitalID not found",
		})
		return
	}

	hospitalClinics, err := c.ClinicService.GetAllHospitalClinic(userHospitalID.(int64))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, hospitalClinics)
}

type HospitalClinicBody struct {
	AddClinicData []string `json:"add_clinic_data"`
}

func (c *ClinicController) CreateHospitalClinic(ctx *gin.Context) {
	var hospitalClinicBody HospitalClinicBody
	session := sessions.Default(ctx)

	if err := ctx.BindJSON(&hospitalClinicBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userRole := session.Get("role")
	if userRole == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "UserRole not found",
		})
		return
	}

	userHospitalID := session.Get("hospital_id")
	if userHospitalID == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "UserHospitalID not found",
		})
		return
	}

	if userRole != "Staff" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Access denied",
		})
		return
	}
	if err := c.ClinicService.CreateHospitalClinic(hospitalClinicBody.AddClinicData, userHospitalID.(int64)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Create successfully",
	})
}

func (c *ClinicController) DeleteHospitalClinic(ctx *gin.Context) {
	clinicName := ctx.Param("clinicName")

	session := sessions.Default(ctx)

	if clinicName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ClinicName parameter is required"})
		return
	}

	userRole := session.Get("role")
	if userRole == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "UserRole not found",
		})
		return
	}

	if userRole != "Staff" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Access denied",
		})
		return
	}

	userHospitalID := session.Get("hospital_id")
	if userHospitalID == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "UserHospitalID not found",
		})
		return
	}

	if err := c.ClinicService.DeleteHospitalClinicByClinicName(userHospitalID.(int64), clinicName); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete successfully",
	})
}

func (c *ClinicController) Update(ctx *gin.Context) {
	var hospitalClinicBody HospitalClinicBody
	session := sessions.Default(ctx)

	if err := ctx.BindJSON(&hospitalClinicBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userRole := session.Get("role")
	if userRole == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "UserRole not found",
		})
		return
	}

	userHospitalID := session.Get("hospital_id")
	if userHospitalID == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "UserHospitalID not found",
		})
		return
	}

	if userRole != "Staff" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Access denied",
		})
		return
	}

	if err := c.ClinicService.Update(hospitalClinicBody.AddClinicData, userHospitalID.(int64)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update successfully",
	})
}
