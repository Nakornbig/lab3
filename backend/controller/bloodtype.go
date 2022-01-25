package controller

import (
	"net/http"

	"github.com/Nakornbig/se-64-main/entity"
	"github.com/gin-gonic/gin"
)

// POST /bloodtypes
func CreateBloodType(c *gin.Context) {
	var bloodtypes entity.Bloodtype
	if err := c.ShouldBindJSON(&bloodtypes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&bloodtypes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bloodtypes})
}

// GET /bloodtype/:id
func GetBloodType(c *gin.Context) {
	var bloodtypes entity.Bloodtype
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM bloodtypes WHERE id = ?", id).Find(&bloodtypes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bloodtypes})
}

// GET /bloodtypes
func ListBloodTypes(c *gin.Context) {
	var bloodtypes []entity.Bloodtype
	if err := entity.DB().Raw("SELECT * FROM bloodtypes").Find(&bloodtypes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bloodtypes})
}

// DELETE /bloodtype/:id
func DeleteBloodType(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM bloodtypes WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bloodtype not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /bloodtypes
func UpdateBloodType(c *gin.Context) {
	var bloodtype entity.Bloodtype
	if err := c.ShouldBindJSON(&bloodtype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", bloodtype.ID).First(&bloodtype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bloodtype not found"})
		return
	}

	if err := entity.DB().Save(&bloodtype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bloodtype})
}
