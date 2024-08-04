package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/MediManage-Backend/cache"
	"github.com/kwa0x2/MediManage-Backend/services"
	"net/http"
)

type JobGroupController struct {
	JobGroupService *services.JobGroupService
	JobGroupCache   *cache.JobGroupCache
}

func (c *JobGroupController) GetAll(ctx *gin.Context) {

	jobGroups, err := c.JobGroupCache.GetAllJobGroups()
	if jobGroups == nil && len(jobGroups) > 0 {
		ctx.JSON(http.StatusOK, jobGroups)
		return
	}

	jobGroups, err = c.JobGroupService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error fetching job groups")
		return
	}

	err = c.JobGroupCache.SetAllJobGroups(jobGroups)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, jobGroups)
}
