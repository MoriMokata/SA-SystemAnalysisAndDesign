package main

import (
	"github.com/MoriMokata/project/backend/controller"
	"github.com/MoriMokata/project/backend/entity"
	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	// User Routes
	r.GET("/api/MedicalTech", controller.GetMedicalTech)
	r.GET("/api/MedicalRecord", controller.ListMedicalRecord)
	r.GET("/api/LabType", controller.ListLabType)
	r.GET("/api/LabRoom", controller.ListLabRoom)
	r.GET("/api/LabResult", controller.ListLabResult)
	r.POST("/api/submit", controller.CreateLabResult)
	
	// Run the server

	r.Run()
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
