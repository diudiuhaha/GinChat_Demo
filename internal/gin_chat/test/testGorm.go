package main

import (
	"ginchat/internal/gin_chat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 连接数据库
	db, err := gorm.Open(mysql.Open("root:123456@tcp(42.192.92.160:13306)/gin_chat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 创建表   /// 如果没有表就自动创建
	err = db.AutoMigrate(&models.UserBasic{})
	if err != nil {
		return
	} // /这里的返回值是一个err ，如果规范写，应该写个错误判断

	// 插入信息到数据库
	userInfo := &models.UserBasic{} // /这里一定是要取*值
	userInfo.Name = "Tom"
	userInfo.Phone = "1234567890"
	db.Create(userInfo) // /当然，每执行一次，就会插入一次

	// 读取数据库信息
	// /根据主键去查找
	println(db.First(&userInfo, 1))

	// 更新值
	// db.Model(userInfo).Update("PassWord", "12345")

	// //删除
	// db.Delete(&userInfo, 1)

}
