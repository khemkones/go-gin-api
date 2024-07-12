package genders

import (
	"fmt"
	"net/http"

	"example.com/sa-65-example/config"
	"example.com/sa-65-example/models"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {

	db := config.DB()

	var genders []models.Genders
	result := db.Find(&genders)

	fmt.Println(result)

	c.JSON(http.StatusOK, &genders)

}
