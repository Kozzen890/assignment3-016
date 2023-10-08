package helpers

import (
	"fmt"
	"net/http"

	"github.com/Kozzen890/assignment3-016/database"
	"github.com/Kozzen890/assignment3-016/models"
	"github.com/gin-gonic/gin"
)

func UpdateDataAPI(c *gin.Context) {
	db, err := database.DbInit()
	if err != nil {
		fmt.Println("Database Error:", err)
		return
	}
	var UptoDate models.WeatherData
	if err := c.ShouldBindJSON(&UptoDate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("Error:", err.Error())
		return
	}

	// pembaruan data ke dalam database
	if err := db.Model(&models.WeatherData{}).Where("id = ?", 1).Updates(&UptoDate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal Update Data"})
		fmt.Println("Error updating data in database:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Terupdate"})
}