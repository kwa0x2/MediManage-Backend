package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/MediManage-Backend/models"
	"github.com/kwa0x2/MediManage-Backend/services"
	"net/http"
)

type UserController struct {
	UserService *services.UserService
}

func (c *UserController) GetAll(ctx *gin.Context) {
	data, err := c.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (c *UserController) Create(ctx *gin.Context) {
	var userBody models.User
	session := sessions.Default(ctx)

	if err := ctx.BindJSON(&userBody); err != nil {
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

	userBody.UserHospitalID = userHospitalID.(int64)
	if err := c.UserService.Create(nil, &userBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Create successfully",
	})
}

type UserBody struct {
	User   *models.User `json:"user"`
	UserID string       `json:"user_id"`
}

func (c *UserController) Update(ctx *gin.Context) {
	var userBody UserBody
	session := sessions.Default(ctx)

	if err := ctx.BindJSON(&userBody); err != nil {
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

	if userRole != "Staff" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Access denied",
		})
		return
	}

	if err := c.UserService.Update(userBody.User, userBody.UserID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update successfully",
	})
}

func (c *UserController) Delete(ctx *gin.Context) {
	userId := ctx.Param("userId")
	session := sessions.Default(ctx)

	if userId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "UserID parameter is required"})
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

	if err := c.UserService.Delete(userId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete successfully",
	})
}
