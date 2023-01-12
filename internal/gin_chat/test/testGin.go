package main

import "github.com/gin-gonic/gin"

func main() {
	///标准的官方demo
	//定义一个路由
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run(":12345")
	if err != nil {
		return
	}

	//发送get请求到 http://127.0.0.1:12345/ping ，会得到 json信息
}
