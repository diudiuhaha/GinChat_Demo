// 用户模块
package models

import (
	"fmt"
	"ginchat/internal/gin_chat/utils"
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
	Identity      string `json:"identity"` ///唯一标识
	ClientIp      string `json:"client_ip"`
	ClientPort    string `json:"client_port"`
	LoginTime     string `json:"login_time"`
	HeartbeatTime string `json:"heartbeat_time"`
	LoginOutTime  string `json:"login_out_time"`
	IsLoginOut    bool   `json:"is_login_out"`
	Device        string `json:"device"`
}

// / 方法1，定义一个方法,等待外部调用
func (table *UserBasic) TableName() string {
	return "user_basic"
}

// /方法2
func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 15)
	utils.DB.Find(&data)

	for _, v := range data {
		fmt.Println(v)
	}
	return data
}
