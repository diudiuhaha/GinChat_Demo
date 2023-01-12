package main

import (
	docs "ginchat/docs"
	"ginchat/internal/gin_chat/router"
	"ginchat/internal/gin_chat/sys_init"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// 初始化
	sys_init.InitConfig()
	sys_init.InitMysql()
	// 路由
	r := router.Router()
	// /swagger
	docs.SwaggerInfo.BasePath = "" // 表示根目录/
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(":12345")
	if err != nil {
		return
	}
}
