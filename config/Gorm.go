package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormConnectWrite() *gorm.DB {
	k := GetMyConfig()
	MysqlHostWrite := k.MYSQL.MYSQL_HOST_WRITE
	MysqlUserWrite := k.MYSQL.MYSQL_USER_WRITE
	MysqlPortWrite := k.MYSQL.MYSQL_PORT_WRITE
	MysqlPassWrite := k.MYSQL.MYSQL_PASS_WRITE
	MysqlDatabaseName := k.MYSQL.MYSQL_DB_NAME
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf(dsn, MysqlUserWrite, MysqlPassWrite, MysqlHostWrite, MysqlPortWrite, MysqlDatabaseName)
	gorm, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	gormDB, _ := gorm.DB()
	gormDB.SetMaxIdleConns(20)
	gormDB.SetConnMaxIdleTime(10)
	gormDB.SetMaxOpenConns(100)
	gormDB.SetConnMaxLifetime(30)
	return gorm
}
