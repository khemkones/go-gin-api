package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/sa-65-example/config"
	"example.com/sa-65-example/models"
)

func GetAll(c *gin.Context) {

	var users []models.Users

	db := config.DB()
	results := db.Preload("Gender").Find(&users)
	if results.Error != nil {
		c.JSON(http.StatusFound, gin.H{"error": results.Error})
		return
	}
	c.JSON(http.StatusCreated, users)

}

func Get(c *gin.Context) {

	ID := c.Param("id")
	var user models.Users

	db := config.DB()
	results := db.Preload("Gender").First(&user, ID)
	if results.Error != nil {
		c.JSON(http.StatusFound, gin.H{"error": results.Error})
		return
	}
	if user.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	c.JSON(http.StatusOK, user)

}

func Update(c *gin.Context) {

	var user models.Users

	UserID := c.Param("id")

	db := config.DB()
	result := db.First(&user, UserID)
	if result.Error != nil {
		c.JSON(http.StatusFound, gin.H{"error": "id not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
}

func Delete(c *gin.Context) {

	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM users WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}
