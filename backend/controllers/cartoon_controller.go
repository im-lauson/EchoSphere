package controllers

import (
	"EchoSphere/backend/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetCartoon(c *gin.Context) {
	// TODO: 实现获取动漫的逻辑

	cartoonData, err := services.GetAllCartoon()
	if err != nil {
		fmt.Printf("获取全部动漫数据:", err)
	}
	c.JSON(200, gin.H{
		"cartoon": cartoonData,
	})

}
