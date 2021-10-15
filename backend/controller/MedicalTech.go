package controller


import (

        "github.com/MoriMokata/project/backend/entity"
        "github.com/gin-gonic/gin"
        "net/http"

)


// GET /MedicalTech/:id

func GetMedicalTech(c *gin.Context) {

	var user entity.MedicalTech

	if err := entity.DB().Table("medical_teches").Where("email = ?", "boss@gmail.com").First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

