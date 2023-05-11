package core

import (
	"fmt"
	"web/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Log.Warnln("未配置pgsql，取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	global.MysqlLog = logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.Log.Fatal(fmt.Sprintf("[%s] mysql连接失败", dsn))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	// fmt.Println("数据库连接成功")
	return db
}


