package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/MediManage-Backend/cache"
	"github.com/kwa0x2/MediManage-Backend/services"
	"net/http"
)

type DistrictController struct {
	DistrictService *services.DistrictService
	DistrictCache   *cache.DistrictCache
}

func (c *DistrictController) GetAll(ctx *gin.Context) {
	data, err := c.DistrictService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error fetching districts")
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (c *DistrictController) GetAllByProvince(ctx *gin.Context) {
	provinceName := ctx.Param("province_name")

	if provinceName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Province parameter is required"})
		return
	}

	districts, err := c.DistrictCache.GetAllDistrictsByProvince(provinceName)
	if districts == nil && len(districts) > 0 {
		ctx.JSON(http.StatusOK, districts)
		return
	}

	districts, err = c.DistrictService.GetAllByProvinceName(provinceName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error fetching districts")
		return
	}

	err = c.DistrictCache.SetAllDistrictsByProvince(districts, provinceName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, districts)
}
