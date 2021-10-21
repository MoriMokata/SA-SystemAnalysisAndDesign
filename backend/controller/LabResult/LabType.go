package controller

import (
	"net/http"
	"github.com/MoriMokata/project/backend/entity"
	"github.com/gin-gonic/gin"
)

func ListLabType(c *gin.Context) {
	var LabType []entity.LabType
	if err := entity.DB().Table("lab_types").Find(&LabType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": LabType})
}