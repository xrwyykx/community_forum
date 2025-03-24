package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

// 全局数据库连接对象
var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		AppConfig.DB.Username,
		AppConfig.DB.Password,
		AppConfig.DB.Host,
		AppConfig.DB.Port,
		AppConfig.DB.DBName,
		AppConfig.DB.Charset,
	)
	var err error
	dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		//将gorm的日志模式为Info，将打印出执行的SQL语句
	})
	if err != nil {
		log.Fatalf("failed to connect DB")
	}
	DB = dbConn
	//log.Println("connect DB successfully")
}
