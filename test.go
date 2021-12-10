package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"log"

	"fmt"

)
type DbLIst struct {
	Isconnect       bool    `json:"isconnect"`
	Id      		int		`json:"Id" form:"Id"`
	Ip				string	`json:"Ip" form:"Ip"`
	Name 			string	`json:"Name" form:"Name"`
	Username 		string	`json:"Username" form:"Username"`
	Password 		string	`json:"Password" form:"Password"`
	Port     		string	`json:"Port"     form:"Port"`
	Service_name  	string	`json:"Service_name" form:"Service_name"`
}



func checkErr(err error) {
	if err != nil {
		println(err)
		//panic(err)

	}
}




func CheckErr(err error) {
	if err != nil {
		log.Println("err",err)
		fmt.Println("err",err)

		//panic(err)
	}
}

var (
	OraSource []DbLIst
)

func main(){
	var err error
	var rows *sql.Rows
	//var sqlstr string
	OraSource = []DbLIst{}


	db,_:=sql.Open("sqlite3", "./database/lixora")

	rows,err = db.Query("select id,name,ip,username,prot,SERVICENAME from dbinfo")
	CheckErr(err)

	for rows.Next() {
		var s  DbLIst
		err = rows.Scan(&s.Id,&s.Name, &s.Ip, &s.Port, &s.Username,&s.Password,&s.Service_name)
		checkErr(err)
		fmt.Print(s)
		OraSource = append(OraSource,s)

	}
	defer rows.Close()
	defer db.Close()
	//return OraSource
	log.Printf("%+v",OraSource)

	fmt.Printf("%+v",OraSource)


}


