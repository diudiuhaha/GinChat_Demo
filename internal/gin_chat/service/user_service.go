// Package service 的方法在route里面进行调用
package service

import (
	"ginchat/internal/gin_chat/models"
	"github.com/gin-gonic/gin"
)

// GetUserInfoList
// @Summary 所有用户
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/GetUserInfoList [get]
func GetUserInfoList(c *gin.Context) {
	data := make([]*models.UserBasic, 15)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUserInfo
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUserInfo [get]
func CreateUserInfo(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	if password != repassword {
		c.JSON(-1, gin.H{
			"message": "两次密码不正确，请重新输入",
		})
	}
	user.PassWord = password
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "新增用户成功",
	})

}
