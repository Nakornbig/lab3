package main

import (
	"github.com/Nakornbig/se-64-main/controller"
	"github.com/Nakornbig/se-64-main/entity"
	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	// BloodType Routes
	r.GET("/bloodtypes", controller.ListBloodTypes)
	r.GET("/bloodtype/:id", controller.GetBloodType)
	r.POST("/bloodtypes", controller.CreateBloodType)
	r.PATCH("/bloodtypes", controller.UpdateBloodType)
	r.DELETE("/bloodtypes/:id", controller.DeleteBloodType)

	// Gender Routes
	r.GET("/genders", controller.ListGenders)
	r.GET("/gender/:id", controller.GetGender)
	r.POST("/genders", controller.CreateGender)
	r.PATCH("/genders", controller.UpdateGender)
	r.DELETE("/genders/:id", controller.DeleteGender)

	// Personnel Routes
	r.GET("/personnels", controller.ListPersonnels)
	r.GET("/personnel/:id", controller.GetPersonnel)
	r.POST("/personnels", controller.CreatePersonnel)
	r.PATCH("/personnels", controller.UpdatePersonnel)
	r.DELETE("/personnels/:id", controller.DeletePersonnel)

	// Prename Routes
	r.GET("/prenames", controller.ListPrenames)
	r.GET("/prename/:id", controller.GetPrename)
	r.POST("/prenames", controller.CreatePrename)
	r.PATCH("/prenames", controller.UpdatePrename)
	r.DELETE("/prenames/:id", controller.DeletePrename)

	// Province Routes
	r.GET("/provinces", controller.ListProvinces)
	r.GET("/province/:id", controller.GetProvince)
	r.POST("/provinces", controller.CreateProvince)
	r.PATCH("/provinces", controller.UpdateProvince)
	r.DELETE("/provinces/:id", controller.DeleteProvince)

	// Patientrecord Routes
	r.GET("/patientrecords", controller.ListPatientrecords)
	r.GET("/patientrecord/:id", controller.GetPatientrecord)
	r.POST("/patientrecords", controller.CreatePatientrecord)
	r.PATCH("/patientrecords", controller.UpdatePatientrecord)
	r.DELETE("/patientrecords/:id", controller.DeletePatientrecord)
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
