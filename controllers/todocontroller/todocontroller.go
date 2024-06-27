package todocontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syauqifut/todowleest/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var todos []models.Todo

	models.DB.Find(&todos)

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func Post(c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := models.DB.Create(&todo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Data created"})
}

func Det(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := models.DB.First(&todo, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func Update(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&todo).Where("id = ?", id).Updates(&todo).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data updated!"})
}

func Del(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if models.DB.Delete(&todo, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data deleted!"})
}
