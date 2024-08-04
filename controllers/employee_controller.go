package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/MediManage-Backend/models"
	"github.com/kwa0x2/MediManage-Backend/services"
	"net/http"
	"strconv"
)

type EmployeeController struct {
	EmployeeService *services.EmployeeService
}

func (c *EmployeeController) GetAll(ctx *gin.Context) {
	session := sessions.Default(ctx)

	userHospitalID := session.Get("hospital_id")
	if userHospitalID == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "UserHospitalID not found",
		})
		return
	}

	employees, err := c.EmployeeService.GetByHospitalId(userHospitalID.(int64))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var response []map[string]interface{}
	for _, emp := range employees {
		workDays := []string{}
		for _, wd := range emp.EmployeeWorkDays {
			workDays = append(workDays, string(wd.Day))
		}
		employeeData := map[string]interface{}{
			"employee_id":              emp.EmployeeID,
			"employee_name":            emp.EmployeeName,
			"employee_surname":         emp.EmployeeSurname,
			"employee_identity_number": emp.EmployeeIdentityNumber,
			"employee_phone_number":    emp.EmployeePhoneNumber,
			"employee_job_group_name":  emp.EmployeeJobGroupName,
			"employee_title_name":      emp.EmployeeTitleName,
			"employee_clinic_name":     emp.EmployeeClinicName,
			"employee_working_days":    workDays,
		}
		response = append(response, employeeData)
	}

	ctx.JSON(http.StatusOK, response)

}

type EmployeeWithWorkDayBody struct {
	Employee            *models.Employee `json:"employee"`
	EmployeeWorkingDays []string         `json:"employee_working_days"`
}

func (c *EmployeeController) CreateWithWorkDay(ctx *gin.Context) {
	var employeeWithWorkDayBody EmployeeWithWorkDayBody
	session := sessions.Default(ctx)

	if err := ctx.BindJSON(&employeeWithWorkDayBody); err != nil {
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

	exists, err := c.EmployeeService.CheckChiefDoctorExists(userHospitalID.(int64))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error checking for chief doctor",
		})
		return
	}

	if exists && employeeWithWorkDayBody.Employee.EmployeeTitleName == "Başhekim" {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": "Hastanede sadece 1 adet başhekim olabilir",
		})
		return
	}

	employeeWithWorkDayBody.Employee.EmployeeHospitalID = userHospitalID.(int64)
	if err := c.EmployeeService.Create(employeeWithWorkDayBody.Employee, employeeWithWorkDayBody.EmployeeWorkingDays); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Create successfully",
	})
}

type EmployeeWithWorkDayUpdateBody struct {
	Employee            *models.Employee `json:"employee"`
	EmployeeID          int64            `json:"employee_id"`
	EmployeeWorkingDays []string         `json:"employee_working_days"`
}

func (c *EmployeeController) Update(ctx *gin.Context) {
	var employeeWithWorkDayUpdateBody EmployeeWithWorkDayUpdateBody
	session := sessions.Default(ctx)

	if err := ctx.BindJSON(&employeeWithWorkDayUpdateBody); err != nil {
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

	exists, err := c.EmployeeService.CheckChiefDoctorExists(userHospitalID.(int64))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error checking for chief doctor",
		})
		return
	}

	if exists && employeeWithWorkDayUpdateBody.Employee.EmployeeTitleName == "Başhekim" {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": "Hastanede sadece 1 adet başhekim olabilir",
		})
		return
	}

	if err := c.EmployeeService.Update(employeeWithWorkDayUpdateBody.Employee, employeeWithWorkDayUpdateBody.EmployeeWorkingDays, employeeWithWorkDayUpdateBody.EmployeeID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update successfully",
	})
}

func (c *EmployeeController) Delete(ctx *gin.Context) {
	employeeIdStr := ctx.Param("employeeId")
	session := sessions.Default(ctx)

	if employeeIdStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "EmployeeID parameter is required"})
		return
	}

	employeeId, err := strconv.ParseInt(employeeIdStr, 10, 64)
	if err != nil {
		fmt.Println("Error converting employeeId:", err)
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

	if err := c.EmployeeService.Delete(employeeId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete successfully",
	})
}
