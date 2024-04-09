package routers

import (
	"EchoSphere/backend/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouters(router *gin.Engine) error {
	var err error

	CartoonRouter := router.Group("/cartoon")
	{
		CartoonRouter.GET("/", controllers.GetCartoon)
	}
	err = router.SetTrustedProxies([]string{"127.0.0.1:7890"})
	if err != nil {
		return err
	}
	err = router.Run(":8080")
	if err != nil {
		return err
	}
	return err
}
