package config

import "github.com/go-sql-driver/mysql"

// 环境配置文件
// 可配置多个环境配置，进行切换

type Env struct {
	Debug bool

	Database     mysql.Config
	MaxIdleConns int
	MaxOpenConns int
	ServerPort   string
	ServerHost   string

	AccessLog     bool
	AccessLogPath string
	ErrorLog      bool
	ErrorLogPath  string
	InfoLog       bool
	InfoLogPath   string

	SqlLog bool

	TemplatePath string // 静态文件相对路径
}

func GetEnv() *Env {
	return &env
}
