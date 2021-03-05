package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 数据库链接
var _conn *gorm.DB

// GetDB 获取数据库链接
func GetDB() *gorm.DB {
	if _conn == nil {
		// dns := fmt.Sprintf(
		// 	"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		// 	config.C.Database.Host,
		// 	config.C.Database.User,
		// 	config.C.Database.Password,
		// 	config.C.Database.Database,
		// 	config.C.Database.Port,
		// )
		dns := "kama:W6rZj7ztkhCEGDf6@tcp(101.200.161.142:3306)/kama?charset=utf8mb4&parseTime=True&loc=Local"
		conn, err := gorm.Open(mysql.Open(dns), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatalln(err)
		}
		_conn = conn
	}
	return _conn
}
