package controller

import (
	"net/http"

	"github.com/Nakornbig/se-64-main/entity"
	"github.com/gin-gonic/gin"
)

// POST /PrenameController
func CreatePrename(c *gin.Context) {
	var prename entity.Prename
	if err := c.ShouldBindJSON(&prename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&prename).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": prename})
}

// GET /prename/:id
func GetPrename(c *gin.Context) {
	var prename entity.Prename
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM return_ways WHERE id = ?", id).Find(&prename).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": prename})
}

// GET /return_ways
func ListPrenames(c *gin.Context) {
	var prenames []entity.Prename
	if err := entity.DB().Raw("SELECT * FROM prenames").Find(&prenames).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": prenames})
}

// DELETE /return_ways/:id
func DeletePrename(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM prename WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prename not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /return_ways
func UpdatePrename(c *gin.Context) {
	var prename entity.Prename
	if err := c.ShouldBindJSON(&prename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", prename.ID).First(&prename); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prename not found"})
		return
	}

	if err := entity.DB().Save(&prename).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": prename})
}
