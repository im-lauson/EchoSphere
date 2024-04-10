package routers

import (
	"EchoSphere/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouters(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"*",
	} // 允许的源地址
	router.Use(cors.New(config))
	CartoonRouter := router.Group("/cartoon")
	{
		CartoonRouter.Any("/", controllers.GetCartoon)
	}
	//router.SetTrustedProxies([]string{"127.0.0.1:7890"})

}
