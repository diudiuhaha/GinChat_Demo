package main

import (
	"ginchat/internal/gin_chat/sys_init"
)

func main() {
	sys_init.InitConfig()
	sys_init.InitMysql()
}
