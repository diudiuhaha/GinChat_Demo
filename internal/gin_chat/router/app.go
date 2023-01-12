package router

import (
	"ginchat/internal/gin_chat/service"
	"github.com/gin-gonic/gin"
)

// / 定义一个路由，调用会返回它的使用方法
func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/index", service.GetIndex)
	r.GET("/user/getUserInfoList", service.GetUserInfoList)
	return r
}
