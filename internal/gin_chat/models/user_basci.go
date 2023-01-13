// 用户模块
package models

import (
	"fmt"
	"ginchat/internal/gin_chat/sys_init"
	"gorm.io/gorm"
)

// 用户模型
// / 设计用户信息的模块，与mysql的字段尽量相同
type UserBasic struct {
	gorm.Model
	Name          string `json:"name"`
	PassWord      string `json:"pass_word"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Identity      string `json:"identity"` // /唯一标识
	ClientIp      string `json:"client_ip"`
	ClientPort    string `json:"client_port"`
	LoginTime     string `json:"login_time"`
	HeartbeatTime string `json:"heartbeat_time"`
	LoginOutTime  string `json:"login_out_time"`
	IsLoginOut    bool   `json:"is_login_out"`
	Device        string `json:"device"`
}

// TableName / 方法1，定义一个方法,等待外部调用
func (table *UserBasic) TableName() string {
	return "user_basic"
}

// GetUserList /方法2用于获取用户信息
func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 15)
	sys_init.DB.Find(&data)

	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

// CreateUser 方法2用于增加用户
func CreateUser(user UserBasic) *gorm.DB {
	return sys_init.DB.Create(&user)
}
