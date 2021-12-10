package config

import (
	"github.com/go-ini/ini"
	"log"
	"strconv"
)

// 本文件建议在代码协同工具(git/svn等)中忽略

var env = Env{
	Debug: true,

	ServerPort: "8081",

	//Database: mysql.Config{
	//	User:                 "root",
	//	Passwd:               "root",
	//	Addr:                 "127.0.0.1:3306",
	//	DBName:               "gin-template",
	//	Collation:            "utf8mb4_unicode_ci",
	//	Net:                  "tcp",
	//	AllowNativePasswords: true,
	//},
	//MaxIdleConns: 50,
	//MaxOpenConns: 100,
	//
	//RedisSessionDb: 1,
	//RedisCacheDb:   2,

	AccessLog:     true,
	AccessLogPath: "storage/logs/access.log",

	ErrorLog:     true,
	ErrorLogPath: "storage/logs/error.log",

	InfoLog:     true,
	InfoLogPath: "storage/logs/info.log",

	TemplatePath: "web/views",

	//APP_SECRET: "YbskZqLNT6TEVLUA9HWdnHmZErypNJpL",
	//AppSecret: "something-very-secret",
}


var (
	Cfg *ini.File

	RUN_MODE   string
	SRV_PORT   int
	SRV_HOST   string

)


func init() {
	var err error
	Cfg, err = ini.Load("config.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'config.ini': %v", err)
	}

	RUN_MODE = Cfg.Section("").Key("RUN_MODE").MustString("")
	if (RUN_MODE==""||RUN_MODE=="release"){
		env.Debug = false
	}

	SRV_PORT = Cfg.Section("").Key("SRV_PORT").MustInt(8188)
	env.ServerPort = strconv.Itoa(SRV_PORT)

	SRV_HOST = Cfg.Section("").Key("SRV_HOST").MustString("0.0.0.0")
	env.ServerHost = SRV_HOST


	println("[ServerHost]:",SRV_HOST)
	println("[ServerPort]:",SRV_PORT)

}
