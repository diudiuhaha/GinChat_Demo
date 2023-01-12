package main

import "ginchat/internal/gin_chat/utils"

func main() {
	utils.InitConfig()
	utils.InitMysql()
}
