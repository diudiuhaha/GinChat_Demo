// int初始化
package sys_init

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB // /一定要放在外面才能被引用

// 初始化配置文件
func InitConfig() {
	viper.SetConfigName("app.yml")
	viper.SetConfigType("yaml")                          // /如果配置文件没有后缀名，则需要声明类型
	viper.AddConfigPath("internal\\gin_chat\\configs\\") // /windows需要反斜杠
	// //读取：
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	// /测试，打印出来

	// //log.Println("configs mysql", viper.Get("mysql"))
	// /log.Println("configs mysql", viper.Get("mysql.dsn"))
}

// InitMysql 初始化mysql
func InitMysql() {
	// 定义日志
	sqlLog := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			// 以下内容可以自动补充
			SlowThreshold:             time.Second, // 慢sql阈值
			Colorful:                  true,        // 颜色
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      false,
			LogLevel:                  logger.Info, // 日志级别
		},
	)
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{Logger: sqlLog})
}
