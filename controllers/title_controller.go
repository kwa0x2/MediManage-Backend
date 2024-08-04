package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/MediManage-Backend/cache"
	"github.com/kwa0x2/MediManage-Backend/services"
	"net/http"
)

type TitleController struct {
	TitleService *services.TitleService
	TitleCache   *cache.TitleCache
}

func (c *TitleController) GetAllTitleByJobGroupName(ctx *gin.Context) {
	jobGroupName := ctx.Param("jobGroupName")

	if jobGroupName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Job group name parameter is required"})
		return
	}

	titles, err := c.TitleCache.GetAllTitleByJobGroupName(jobGroupName)
	if titles == nil && len(titles) > 0 {
		ctx.JSON(http.StatusOK, titles)
		return
	}

	titles, err = c.TitleService.GetAllByJobGroupName(jobGroupName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error fetching titles")
		return
	}

	err = c.TitleCache.SetAllTitleByJobGroupName(titles, jobGroupName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, titles)
}
