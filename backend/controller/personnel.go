package controller

import (
	"net/http"

	"github.com/Nakornbig/se-64-main/entity"
	"github.com/gin-gonic/gin"
)

// POST /MedicalRecordStaffController
func CreatePersonnel(c *gin.Context) {
	var personnel entity.Personnel
	if err := c.ShouldBindJSON(&personnel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&personnel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": personnel})
}

// GET /medicalrecord/:id
func GetPersonnel(c *gin.Context) {
	var personnel entity.Personnel
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM Personnel WHERE id = ?", id).Find(&personnel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": personnel})
}

// GET /MedicalRecordStaffController
func ListPersonnels(c *gin.Context) {
	var personnels []entity.Personnel
	if err := entity.DB().Raw("SELECT * FROM personnels").Find(&personnels).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": personnels})
}

// DELETE /MedicalRecordStaffController/:id
func DeletePersonnel(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM Personnel WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicalrecord not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /MedicalRecordStaffController
func UpdatePersonnel(c *gin.Context) {
	var personnel entity.Personnel
	if err := c.ShouldBindJSON(&personnel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", personnel.ID).First(&personnel); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "personnel not found"})
		return
	}

	if err := entity.DB().Save(&personnel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": personnel})
}
