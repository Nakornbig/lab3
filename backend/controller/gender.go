package controller

import (
	"net/http"

	"github.com/Nakornbig/se-64-main/entity"
	"github.com/gin-gonic/gin"
)

// POST /GenderController
func CreateGender(c *gin.Context) {
	var genders entity.Gender
	if err := c.ShouldBindJSON(&genders); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&genders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": genders})
}

// GET /gender/:id
func GetGender(c *gin.Context) {
	var genders entity.Gender
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM genders WHERE id = ?", id).Find(&genders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": genders})
}

// GET /genders
func ListGenders(c *gin.Context) {
	var genders []entity.Gender
	if err := entity.DB().Raw("SELECT * FROM genders").Find(&genders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": genders})
}

// DELETE /GenderController/:id
func DeleteGender(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM genders WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /GenderController
func UpdateGender(c *gin.Context) {
	var genders entity.Gender
	if err := c.ShouldBindJSON(&genders); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", genders.ID).First(&genders); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}

	if err := entity.DB().Save(&genders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": genders})
}
