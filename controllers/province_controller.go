package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/MediManage-Backend/cache"
	"github.com/kwa0x2/MediManage-Backend/services"
	"net/http"
)

type ProvinceController struct {
	ProvinceService *services.ProvinceService
	ProvinceCache   *cache.ProvinceCache
}

func (c *ProvinceController) GetAll(ctx *gin.Context) {

	provinces, err := c.ProvinceCache.GetAllProvince()
	if provinces == nil && len(provinces) > 0 {
		ctx.JSON(http.StatusOK, provinces)
		return
	}

	provinces, err = c.ProvinceService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error fetching provinces")
		return
	}

	err = c.ProvinceCache.SetAllProvince(provinces)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, provinces)
}
