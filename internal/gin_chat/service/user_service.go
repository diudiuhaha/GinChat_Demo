// / service的方法在route里面进行调用
package service

import (
	"ginchat/internal/gin_chat/models"
	"github.com/gin-gonic/gin"
)

// / 获取人员信息的列表的方法
func GetUserInfoList(c *gin.Context) {
	data := make([]*models.UserBasic, 15)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}