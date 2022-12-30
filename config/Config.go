package config

import (
	"fmt"

	"dot.go/helper"
	"github.com/spf13/viper"
)

type MyConfig struct {
	APP   AppConfig
	REDIS RedisConfig
	MYSQL MySQLConfig
}

type AppConfig struct {
	API_KEY            string
	APP_NAME           string
	APP_DOMAIN         string
	MY_PACKAGE         string
	MY_PORT            string
	MY_URL             string
	VERSION            string
	SHOW_LOG           string
	IS_PRODUCTION      bool
	JWT_SECRET         string
	JWT_REFRESH_SECRET string
	MAINTENANCE        bool
	MSG_MAINTENANCE    string
}

type MySQLConfig struct {
	MYSQL_HOST_WRITE string
	MYSQL_USER_WRITE string
	MYSQL_PASS_WRITE string
	MYSQL_PORT_WRITE string
	MYSQL_DB_NAME    string
}

type RedisConfig struct {
	REDIS_HOST_LOCAL string
	REDIS_PORT_LOCAL string
	REDIS_PASS_LOCAL string
}

func GetMyConfig() MyConfig {
	// Set the file name of the configurations file
	funcNow := "GetMyConfig"
	viper.SetConfigName("xconfig")
	// Set the path to look for the configurations file
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/.conf.d/.dot_test/")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	var myconfig MyConfig
	if err := viper.ReadInConfig(); err != nil {
		helper.LogError(err.Error(), "func:"+funcNow, "script:load config from yaml")
		fmt.Printf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&myconfig)
	if err != nil {
		helper.LogError(err.Error(), "func:"+funcNow, "script:Unmarshal config from yaml")
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	return myconfig
}
