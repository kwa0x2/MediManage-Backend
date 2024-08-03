package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/MediManage-Backend/models"
	"github.com/kwa0x2/MediManage-Backend/services"
	"github.com/kwa0x2/MediManage-Backend/utils"
	"gorm.io/gorm"
	"net/http"
)

type AuthController struct {
	AuthService *services.AuthService
	UserService *services.UserService
}

type RegisterBody struct {
	Hospital *models.Hospital `json:"hospital"`
	User     *models.User     `json:"user"`
}

func (c *AuthController) Register(ctx *gin.Context) {
	var registerBody RegisterBody

	if err := ctx.BindJSON(&registerBody); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, _ := utils.HashPassword(registerBody.User.UserPassword)
	registerBody.User.UserPassword = string(hashedPassword)

	if err := c.AuthService.Register(registerBody.Hospital, registerBody.User); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Kayıt başarılı",
	})
}

type LoginBody struct {
	UserIdentifier string `json:"user_identifier"` // email or phone number
	UserPassword   string `json:"user_password"`
}

func (c *AuthController) Login(ctx *gin.Context) {
	var loginBody LoginBody

	if err := ctx.BindJSON(&loginBody); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data, err := c.UserService.GetUserPasswordByIdentifier(loginBody.UserIdentifier)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid credentials",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !utils.CheckPassword(loginBody.UserPassword, data.UserPassword) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	session := sessions.Default(ctx)
	session.Set("uuid", data.UserID.String())
	session.Set("role", string(data.UserRole))
	err = session.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println(session.ID())

	ctx.JSON(http.StatusOK, gin.H{
		"uuid": data.UserID.String(),
	})
}
