package main

import (
	"EchoSphere/backend/config"
	"EchoSphere/backend/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	InitConfig := config.InitConfig{}
	InitConfig.InitLogger()
	InitConfig.InitDB()
	InitConfig.InitRedis()

	// 初始化GIN引擎
	router := gin.Default()
	// 初始化路由
	routers.InitRouters(router)

	// 启动服务

}
