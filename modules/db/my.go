package db
//
//import (
//	"fmt"
//	"github.com/jinzhu/gorm"
//)
//
////定义db连接信息
//
//type mysqlModel struct {
//	Host        string `yaml:"host"`
//	Port        int    `yaml:"port"`
//	User        string `yaml:"user"`
//	Password    string `yaml:"password"`
//	Database    string `yaml:"database"`
//	MaxIdLeConn int    `yaml:"maxidleconn"`
//	MaxOpenConn int    `yaml:"maxopenconn"`
//	Debug       bool   `yaml:"debug"`
//	IsPlural    bool   `yaml:"isplural"`
//	TablePrefix string `yaml:"tableprefix"`
//}
//
//var DB *gorm.DB
////var mysqlConfig = mysqlModel{
////	Host:        "127.0.0.1",
////	Port:        3308,
////	User:        "root",
////	Password:    "root123456",
////	Database:    "data_center",
////	MaxIdLeConn: 10,
////	MaxOpenConn: 10,
////	Debug:       false,
////	IsPlural:    true,
////	TablePrefix: "data_center_",
////}
//
//func MysqlInit(mysqlConfig) *gorm.DB {
//	var err error
//	DbUrl := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8",
//		mysqlConfig.User,
//		mysqlConfig.Password,
//		mysqlConfig.Host,
//		mysqlConfig.Port,
//		mysqlConfig.Database)
//	DB, err = gorm.Open("mysql", DbUrl)
//	if err != nil {
//		fmt.Printf("mysql connect error %v", err)
//	}
//	if DB.Error != nil {
//		fmt.Printf("database error %v", DB.Error)
//	}
//	//最大打开的连接数
//	DB.DB().SetMaxOpenConns(mysqlConfig.MaxOpenConn)
//	//最大空闲连接数
//	DB.DB().SetMaxIdleConns(mysqlConfig.MaxIdLeConn)
//	//允许表名复数
//	DB.SingularTable(mysqlConfig.IsPlural)
//	DB.LogMode(mysqlConfig.Debug)
//
//	return DB
//}