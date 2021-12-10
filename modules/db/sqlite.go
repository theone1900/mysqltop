package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//"time"
)


type Dbinfo struct {
	Id      		int		`gorm:"AUTO_INCREMENT" json:"Id" form:"Id"`
	Name 			string	`gorm:"not null" json:"Name"     form:"Name"`
	Ip				string	`gorm:"not null" json:"Ip"       form:"Ip"`
	Username 		string	`gorm:"not null" json:"Username" form:"Username"`
	Password 		string	`gorm:"not null" json:"Password" form:"Password"`
	Port     		string	`gorm:"not null" json:"Port"     form:"Port"`
	Service_name  	string	`gorm:"not null" json:"Service_name" form:"Service_name"`
}


func InitDb() *gorm.DB {
		// Openning file
		db, err := gorm.Open("sqlite3", "./database/lixora")
		db.LogMode(true)
		// Error
		if err != nil {
		panic(err)
	}
		// Creating the table
		if !db.HasTable(&Dbinfo{}) {
		db.CreateTable(&Dbinfo{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Dbinfo{})
	}

		return db
	}



