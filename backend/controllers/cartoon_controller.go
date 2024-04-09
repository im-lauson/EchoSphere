package controllers

import (
	"EchoSphere/backend/services"
	"github.com/gin-gonic/gin"
)

func GetCartoon(c *gin.Context) {
	// TODO: 实现获取动漫的逻辑
	cartoon := services.CartoonService{}
	c.JSON(200, gin.H{
		"code": 0,
		"data": cartoon,
	})

}
