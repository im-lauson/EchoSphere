package routers

import (
	"EchoSphere/backend/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouters(router *gin.Engine) {

	CartoonRouter := router.Group("/cartoon")
	{
		CartoonRouter.Any("/", controllers.GetCartoon)
	}
	//router.SetTrustedProxies([]string{"127.0.0.1:7890"})

}
