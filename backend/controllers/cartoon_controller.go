package controllers

import (
	"EchoSphere/backend/services"
	"github.com/gin-gonic/gin"
)

func GetCartoon(c *gin.Context) {
	// TODO: 实现获取动漫的逻辑
	app := services.CartoonService{}
	cartoonData, err := app.GetAllCartoon()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"cartoon": cartoonData,
	})

}
