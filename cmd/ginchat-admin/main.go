package main

import (
	"ginchat/internal/gin_chat/router"
	"ginchat/internal/gin_chat/utils"
)

func main() {
	//初始化
	utils.InitConfig()
	utils.InitMysql()
	//路由
	r := router.Router()
	err := r.Run(":12345")
	if err != nil {
		return
	}
}
