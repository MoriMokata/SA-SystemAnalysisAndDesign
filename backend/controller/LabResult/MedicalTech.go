package controller


import (

        "github.com/MoriMokata/project/backend/entity"
        "github.com/gin-gonic/gin"
        "net/http"

)


// GET /MedicalTech/:id

func GetMedicalTech(c *gin.Context) {
	var user entity.MedicalTech
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM medical_teches WHERE id = ?", id).Scan(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

