package controller

import (
	"github.com/MoriMokata/project/backend/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// GET /Screenings
func PreloadScreenings(c *gin.Context) {
	var Screenings []entity.Screening
	if err := entity.DB().Preload("Nurse").Preload("MedRec").Preload("Disease").Raw("SELECT * FROM screenings").Find(&Screenings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Screenings})
}
