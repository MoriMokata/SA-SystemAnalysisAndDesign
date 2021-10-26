package main

import (
	"github.com/MoriMokata/project/backend/controller/LabResult"
	"github.com/MoriMokata/project/backend/entity"
	"github.com/MoriMokata/project/backend/middlewares"
	"github.com/gin-gonic/gin"
	
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())
	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			protected.GET("/api/MedicalTech/:id", controller.GetMedicalTech)
			protected.GET("/api/MedicalRecord", controller.ListMedicalRecord)
			protected.GET("/api/LabType", controller.ListLabType)
			protected.GET("/api/LabRoom", controller.ListLabRoom)
			protected.GET("/api/LabResult", controller.ListLabResult)
			protected.POST("/api/submit", controller.CreateLabResult)	

		}
	}
	r.POST("/api/login", controller.Login)
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
