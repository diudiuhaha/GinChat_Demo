// int初始化
package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB ///一定要放在外面才能被引用

// 初始化配置文件
func InitConfig() {
	viper.SetConfigName("app.yml")
	viper.SetConfigType("yaml")                         ///如果配置文件没有后缀名，则需要声明类型
	viper.AddConfigPath("internal\\gin_chat\\config\\") ///windows需要反斜杠
	////读取：
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	///测试，打印出来

	////log.Println("config mysql", viper.Get("mysql"))
	///log.Println("config mysql", viper.Get("mysql.dsn"))
}

// 初始化mysql
func InitMysql() {

	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{})
}
