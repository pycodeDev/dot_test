package config

import (
	"database/sql"
	"time"
)

func MySQLConnect() *sql.DB {
	k := GetMyConfig()
	MysqlHostWrite := k.MYSQL.MYSQL_HOST_WRITE
	MysqlUserWrite := k.MYSQL.MYSQL_USER_WRITE
	MysqlPortWrite := k.MYSQL.MYSQL_PORT_WRITE
	MysqlPassWrite := k.MYSQL.MYSQL_PASS_WRITE
	MysqlDatabaseName := k.MYSQL.MYSQL_DB_NAME
	//fmt.Println(MysqlHostWrite,MysqlUserWrite,MysqlPassWrite,MysqlPortWrite,MysqlDatabaseName);
	db_write, _ := sql.Open("mysql", MysqlUserWrite+":"+MysqlPassWrite+"@tcp("+MysqlHostWrite+":"+MysqlPortWrite+")/"+MysqlDatabaseName)
	db_write.SetMaxOpenConns(100)
	db_write.SetMaxIdleConns(50)
	db_write.SetConnMaxIdleTime(time.Second * 10)
	db_write.SetConnMaxLifetime(time.Second * 20)
	return db_write
}
