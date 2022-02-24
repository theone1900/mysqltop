package v1

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	//"database/sql"
	"github.com/gin-gonic/gin"
	"mysqltop/modules/db"
	"net/http"
)

type DBHang struct {
	Isconnect bool
	//DBS    conn.OraSourceConfig
	DB *sql.DB
}

////
type Dbinfo struct {
	Id           int    `json:"Id" form:"Id"`
	Ip           string `json:"Ip" form:"Ip"`
	Name         string `json:"Name" form:"Name"`
	Username     string `json:"Username" form:"Username"`
	Password     string `json:"Password" form:"Password"`
	Port         string `json:"Port"     form:"Port"`
	Service_name string `json:"Service_name" form:"Service_name"`
	Dbtype       string `json:"Dbtype" form:"Dbtype"`
}

type MysqlTopsess struct {
	Conn_id           int    `json:"conn_id" form:"conn_id"`
	User              string `json:"user" form:"user"`
	Db                string `json:"db" form:"db"`
	Command           string `json:"command" form:"command"`
	State             string `json:"state" form:"state"`
	Time              string `json:"time"     form:"time"`
	Current_statement string `json:"current_statement" form:"current_statement"`
	Last_statement    string `json:"last_statement" form:"last_statement"`
	Program_name      string `json:"program_name" form:"program_name"`
}

type MysqlTopsql struct {
	SCHEMA_NAME       string `json:"SCHEMA_NAME" form:"SCHEMA_NAME"`
	DIGEST_TEXT       string `json:"DIGEST_TEXT" form:"DIGEST_TEXT"`
	COUNT_STAR        string `json:"COUNT_STAR" form:"COUNT_STAR"`
	Sum_time          string `json:"sum_time" form:"sum_time"`
	Min_time          string `json:"min_time" form:"min_time"`
	Avg_time          string `json:"avg_time"     form:"avg_time"`
	Max_time          string `json:"max_time" form:"max_time"`
	SUM_LOCK_TIME     string `json:"SUM_LOCK_TIME" form:"SUM_LOCK_TIME"`
	SUM_ROWS_AFFECTED string `json:"SUM_ROWS_AFFECTED" form:"SUM_ROWS_AFFECTED"`
	SUM_ROWS_SENT     string `json:"SUM_ROWS_SENT" form:"SUM_ROWS_SENT"`
	SUM_ROWS_EXAMINED string `json:"SUM_ROWS_EXAMINED " form:"SUM_ROWS_EXAMINED "`
}

type MysqlDSN struct {
	Ip       string
	Username string
	Password string
	Port     int
}

type Ret struct {
}

//
//var DBH []DBHang
//
//func init(){
//	//conn.GetDBSource(0,0,"")
//
//	//fmt.Println("OraSource=",conn.OraSource)
//
//	//for _,v:=range conn.OraSource{
//	//	dbh:= DBHang{0,0,false,v,nil}
//	//	oradb,err:=m.InitDB(v)
//	//	if err!=nil{
//	//		dbh.Isconnect=false
//	//		dbh.DB=nil
//	//	}else{
//	//		dbh.Isconnect=true
//	//		dbh.DB=oradb
//	//	}
//	//	DBH = append(DBH,dbh)
//	//}
//	//DiscoveryHangJob()
//}
//
//
//func Core(c *gin.Context){
//
//	c.HTML(200,"index.html",gin.H{
//		"title":"index",
//		"statusCode":0,
//	})
//}
//
//func Default(c *gin.Context){
//	c.Redirect(http.StatusTemporaryRedirect,"/login")
//}
//
//func TestIFrame(c *gin.Context){
//	//c.Writer.Header().Set("x-frame-options","DENY")
//
//	c.HTML(http.StatusOK,"1.html",gin.H{
//		"title":      "index",
//		"copyright":  COPYRIGHT,
//		"statusCode": 0,
//	})
//}
//
//func ListPage(c *gin.Context){
//
//	c.HTML(http.StatusOK,"index.html",gin.H{
//		"title":      "index",
//		"copyright":  COPYRIGHT,
//		"statusCode": 0,
//	})
//}
//

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})

}

//api：数据接口：获取后台数据库中配置的数据库连接信息，返回json 格式数据
func GetUsers(c *gin.Context) {
	// Connection to the database
	db := db.InitDb()
	// Close connection database
	defer db.Close()

	var dbinfo []Dbinfo
	// SELECT * FROM dbinfos
	db.Find(&dbinfo)

	// Display JSON result
	c.JSON(http.StatusOK, gin.H{"code": 0,
		"msg":   "",
		"count": 1000,
		"data":  dbinfo,
	})

}

//api：数据接口
//mysql top session
//获取对应mysql top session 信息，返回json 格式数据

func Getmysqltopsess(c *gin.Context) {

	//c.ShouldBind(&id)
	ID := c.PostForm("id")

	//print id
	fmt.Println("Id：", ID)

	//查询数据
	//data := &[]*MysqlDSN{} //定义结果集数组
	//db.Table("dbinfos").Where("id = ?",ID).Find(data)
	data := []MysqlDSN{}

	// Connection to the local database
	db := db.InitDb()
	// Close connection database
	//defer db.Close()

	//var dbinfo []Dbinfo
	// SELECT * FROM dbinfos
	//db.Find(&dbinfo)
	//db.Find(data)

	db.Table("dbinfos").Where("id = ?", ID).Find(&data)
	fmt.Printf("x 的数据类型是:  ", data[0].Ip, data[0].Username, data[0].Password, data[0].Port)

	// Connection to the remote mysql db
	DbUrl := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8",
		data[0].Username,
		data[0].Password,
		data[0].Ip,
		data[0].Port,
		"sys")
	fmt.Println("DbUrl:  ", DbUrl)

	var mydb *gorm.DB

	mydb, err := gorm.Open("mysql", DbUrl)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	defer mydb.Close()
	//最大打开的连接数
	mydb.DB().SetMaxOpenConns(5)
	//最大空闲连接数
	mydb.DB().SetMaxIdleConns(10)
	defer db.Close()

	ret := []MysqlTopsess{}
	mydb.Table("x$session").Where("state != ?", "Sleep").Find(&ret)

	fmt.Printf("ret 的数据类型是:  ", ret[0])

	// Close connection database

	// Display JSON result
	c.JSON(http.StatusOK, gin.H{"code": 0,
		"msg":   "",
		"count": 1000,
		"data":  ret,
	})

}

//mysql top sql
//获取对应mysql top sql 信息，返回json 格式数据

//mysql Db数据库连接池
var DB *sql.DB

//注意方法名大写，就是public,mysql 数据库初始化
func InitmyDB(dsn string) {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	//dsn := strings.Join([]string{dsn.Username, ":", dsn.Password, "@tcp(", dsn.Ip, ":", dsn.Port, ")/", '', "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", dsn)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(10)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(5)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("[db connet err]open database fail", err)
		return
	}
	fmt.Println("[db connet succes]connnect success")
}

//查询操作demo
//func Query() {
//	var user User
//	rows, e := DB.Query("select * from user where id in (1,2,3)")
//	if e == nil {
//		errors.New("query incur error")
//	}
//	for rows.Next() {
//		e := rows.Scan(user.sex, user.phone, user.name, user.id, user.age)
//		if e != nil {
//			fmt.Println(json.Marshal(user))
//		}
//	}
//	rows.Close()
//	DB.QueryRow("select * from user where id=1").Scan(user.age, user.id, user.name, user.phone, user.sex)
//
//	stmt, e := DB.Prepare("select * from user where id=?")
//	query, e := stmt.Query(1)
//	query.Scan()
//}

func Getmysqltopsql(c *gin.Context) {
	//c.ShouldBind(&id)
	ID := c.PostForm("id")
	//print id
	fmt.Println("Id：", ID)

	//查询数据
	//data := &[]*MysqlDSN{} //定义结果集数组
	//db.Table("dbinfos").Where("id = ?",ID).Find(data)
	data := []MysqlDSN{}

	// Connection to the local database
	db := db.InitDb()
	// Close connection database
	//defer db.Close()

	//var dbinfo []Dbinfo
	// SELECT * FROM dbinfos
	//db.Find(&dbinfo)
	//db.Find(data)

	db.Table("dbinfos").Where("id = ?", ID).Find(&data)
	fmt.Println("[数据库连接信息]:  ", data[0].Ip, data[0].Username, data[0].Password, data[0].Port)

	// Connection to the remote mysql db
	DbUrl := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8",
		data[0].Username,
		data[0].Password,
		data[0].Ip,
		data[0].Port,
		"performance_schema")
	fmt.Println("[DbUrl]: ", DbUrl)

	InitmyDB(DbUrl)

	var mysqltop MysqlTopsql
	var ret []MysqlTopsql
	rows, e := DB.Query("SELECT SCHEMA_NAME, DIGEST_TEXT, COUNT_STAR, sys.format_time ( SUM_TIMER_WAIT ) AS sum_time, sys.format_time ( MIN_TIMER_WAIT ) AS min_time, sys.format_time ( AVG_TIMER_WAIT ) AS avg_time, sys.format_time ( MAX_TIMER_WAIT ) AS max_time, sys.format_time ( SUM_LOCK_TIME ) AS SUM_LOCK_TIME, SUM_ROWS_AFFECTED, SUM_ROWS_SENT, SUM_ROWS_EXAMINED  FROM performance_schema.events_statements_summary_by_digest  WHERE SCHEMA_NAME IS NOT NULL  ORDER BY COUNT_STAR DESC")
	fmt.Println("mysql db query succes")
	if e != nil {
		fmt.Println("[DB.Query error]:", e)
	}
	for rows.Next() {
		//fmt.Println("1--row start")
		e := rows.Scan(&mysqltop.SCHEMA_NAME, &mysqltop.DIGEST_TEXT, &mysqltop.COUNT_STAR, &mysqltop.Sum_time, &mysqltop.Min_time, &mysqltop.Avg_time, &mysqltop.Max_time, &mysqltop.SUM_LOCK_TIME, &mysqltop.SUM_ROWS_AFFECTED, &mysqltop.SUM_ROWS_SENT, &mysqltop.SUM_ROWS_EXAMINED)
		//fmt.Println("2--row end")

		if e != nil {
			fmt.Println(json.Marshal(mysqltop))
		}

		ret = append(ret, mysqltop)
	}
	rows.Close()
	fmt.Println(ret)

	//var mydb *gorm.DB
	//
	//mydb, err := gorm.Open("mysql", DbUrl)
	//if err != nil {
	//	fmt.Printf("[DB connet error]mysql connect error %v", err)
	//}
	//defer mydb.Close()
	////最大打开的连接数
	//mydb.DB().SetMaxOpenConns(5)
	////最大空闲连接数
	//mydb.DB().SetMaxIdleConns(10)
	//defer db.Close()
	//
	//ret1 := []MysqlTopsql{}
	//mydb.Table("events_statements_summary_by_digest").Where("SCHEMA_NAME is not null").Find(&ret1)
	//
	//fmt.Println("ret1 的数据类型是:  ", ret1[0])

	// Close connection database

	// Display JSON result
	c.JSON(http.StatusOK, gin.H{"code": 0,
		"msg":   "",
		"count": 1000,
		"data":  ret,
	})

}

//api：删除后台数据库中配置的数据库连接信息
func Postdeldblist(c *gin.Context) {

	//c.ShouldBind(&id)
	ID := c.PostForm("id")

	// Connection to the database
	db := db.InitDb()
	// Close connection database
	defer db.Close()

	var dbinfo []Dbinfo

	db.Where("id = ?", ID).Delete(&dbinfo)

	// Display JSON result
	c.JSON(http.StatusOK, gin.H{"code": 0,
		"msg":        "del dbinfo succes",
		"statusCode": "0",
	})
}

//api：添加后台数据库中配置的数据库连接信息
func Postadddblist(c *gin.Context) {

	//c.ShouldBind(&id)
	var dbinfo Dbinfo

	// ---> 绑定数据
	c.ShouldBindJSON(&dbinfo)

	// Connection to the database
	db := db.InitDb()
	// Close connection database
	defer db.Close()

	db.Create(&dbinfo)

	// Display JSON result
	c.JSON(http.StatusOK, gin.H{
		"msg":  "add dbinfo succes",
		"data": "success",
	})

}

func Postupdatedblist(c *gin.Context) {
	//c.ShouldBind(&id)
	//ID := c.PostForm("id")
	//NAME :=c.PostForm("name")
	//USERNAME := c.PostForm("username")
	//PASSWORD := c.PostForm("password")
	//PORT := c.PostForm("port")
	//SERVICE_NAME := c.PostForm("service_name")
	//DBTYPE := c.PostForm("dbtype")
	var dbinfo []Dbinfo
	c.ShouldBindJSON(&dbinfo)

	// Connection to the database
	db := db.InitDb()
	// Close connection database
	defer db.Close()

	db.Save(&dbinfo)

	// Display JSON result
	c.JSON(http.StatusOK, gin.H{"code": 0,
		"msg":   "update dbinfo succes",
		"count": 1000,
		//"data":  ,
	})
}

func Finddblist(c *gin.Context) {
	// Display JSON result
	c.JSON(http.StatusOK, gin.H{"code": 0,
		"msg":   "find dbinfo succes",
		"count": 1000,
		//"data":  ,
	})

}

func CheckErr(err error) {
	if err != nil {
		log.Println("err", err)
		fmt.Println("err", err)

		panic(err)
	}
}

var (
	OraSource []Dbinfo
)

//############################################################################
//页面展现api函数
func Dblist(c *gin.Context) {
	//Getdbinfo()
	c.HTML(http.StatusOK, "db_list.html", gin.H{
		"code":  0,
		"msg":   "",
		"count": 1000,
		"data":  OraSource,
	})

}

func Topprocesslist(c *gin.Context) {
	//Getdbinfo()
	c.HTML(http.StatusOK, "db_list_topprocess.html", gin.H{
		"code":  0,
		"msg":   "",
		"count": 1000,
		"data":  OraSource,
	})

}

func Healthcheck(c *gin.Context) {
	//Getdbinfo()
	c.HTML(http.StatusOK, "db_list_healcheck.html", gin.H{
		"code":  0,
		"msg":   "",
		"count": 1000,
		"data":  OraSource,
	})

}

func Returnhcreport(c *gin.Context) {
	//Getdbinfo()
	c.HTML(http.StatusOK, "db_health_mysql_check_20210511112526.html", nil)

}

func Seccheck(c *gin.Context) {
	//Getdbinfo()
	c.HTML(http.StatusOK, "db_list_seccheck.html", gin.H{
		"code":  0,
		"msg":   "",
		"count": 1000,
		"data":  OraSource,
	})

}

func Dblistadd(c *gin.Context) {
	c.HTML(http.StatusOK, "add_dbinfo.html", gin.H{})

}

func Mysqltopsess(c *gin.Context) {
	//获取跳转页面数据行ID值
	ID := c.PostForm("id")
	c.HTML(http.StatusOK, "db_list_mysqltopsess.html", gin.H{
		"id": ID, //传递ID值到db_list_mysqltopsess.html页面
	})
}

//mysql 慢sql 分析页面汇总
func Mysqlslowsqllist(c *gin.Context) {
	//获取跳转页面数据行ID值
	ID := c.PostForm("id")
	c.HTML(http.StatusOK, "db_list_topsql.html", gin.H{
		"id": ID, //传递ID值到db_list_mysqlslowsql.html页面
	})
}

//mysql 慢sql 分析页面
func Mysqlslowsql(c *gin.Context) {
	//获取跳转页面数据行ID值
	ID := c.PostForm("id")
	c.HTML(http.StatusOK, "db_list_mysqlslowsql_list.html", gin.H{
		"id": ID, //传递ID值到db_list_mysqlslowsql.html页面
	})
}

func Dblistedit(c *gin.Context) {
	ID := c.PostForm("id")
	c.HTML(http.StatusOK, "edit_dbinfo.html", gin.H{
		"id": ID,
	})
}

func Dbinfo1(c *gin.Context) {
	//c.String(http.StatusOK, "some string")
	c.HTML(http.StatusOK, "line.html", gin.H{})
}

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{
		//"copyright":  "Powered by 黄林杰",
	})
}

//func Dbinfo(c *gin.Context){
//	c.String(http.StatusOK, "some string")
//}

//
//func HAHAHAHA(c *gin.Context){
//
//	c.HTML(http.StatusOK,"hahahahaha.html",gin.H{
//		"title":      "indexes",
//		"copyright":  COPYRIGHT,
//		"statusCode": 0,
//	})
//}
//
//func ListOraConnector(c *gin.Context){
//
//	var currnet_page int
//
//	ipaddr := c.Query("Ip")
//	page := ""
//	pagesize := 10
//	if (page==""){
//		currnet_page = 1
//	}else{
//		currnet_page,_ = strconv.Atoi(page)
//	}
//	if (ipaddr == ""){
//		pagesize = 0
//	}
//
//	//conn.GetDBSource(currnet_page,pagesize ,ipaddr)
//
//	c.JSON(http.StatusOK,gin.H{
//		"status":"ok",
//		"data":conn.OraSource,
//		"count":len(conn.OraSource),
//		"page":currnet_page,
//		"pagesize":pagesize,
//		"statusCode":0,
//	})
//}
//
//func ListOraConnector2(c *gin.Context){
//
//	var currnet_page int
//
//	ipaddr := c.Query("Ip")
//	page := ""
//	pagesize := 0
//	if (page==""){
//		currnet_page = 1
//	}else{
//		currnet_page,_ = strconv.Atoi(page)
//	}
//	if (ipaddr == ""){
//		pagesize = 0
//	}
//
//	conn.GetDBSource(currnet_page,pagesize ,ipaddr)
//
//	var hanglist []Hanglist
//	for _,v:=range conn.OraSource{
//		var h Hanglist
//		for n,p:=range DBH {
//			if v.Id==p.DBS.Id{
//				h.Waiters= DBH[n].Waiters
//				h.Holders= DBH[n].Holders
//				h.Isconnect= DBH[n].Isconnect
//
//			}
//		}
//		h.Id=v.Id
//		h.Name=v.Name
//		h.Service_name=v.Service_name
//		h.Username=v.Username
//		h.Password=v.Password
//		h.Port=v.Port
//		h.Ip=v.Ip
//		hanglist=append(hanglist,h)
//		//println("hanglist :",h)
//	}
//	c.JSON(http.StatusOK,gin.H{
//		"status":"ok",
//		"data":hanglist,
//		"count":len(conn.OraSource),
//		"page":currnet_page,
//		"pagesize":pagesize,
//		"statusCode":0,
//	})
//}
//
//func LoginfoPage(c *gin.Context){
//	c.HTML(http.StatusOK,"killlog.html",gin.H{
//
//	})
//}
//
//func AddOraConnPage(c *gin.Context){
//	c.HTML(http.StatusOK,"addora.html",gin.H{
//
//	})
//}
//
//func SaveOraConnPage(c *gin.Context){
//	c.HTML(http.StatusOK,"saveora.html",gin.H{
//
//	})
//}
//
//func TryConnectin(c *gin.Context){
//	ipaddr := c.PostForm("ipaddr")
//	port := c.PostForm("port")
//
//	srvname := c.PostForm("srvname")
//	username := c.PostForm("username")
//	password:= c.PostForm("password")
//
//
//	_,err:= strconv.Atoi(port)
//	if err!=nil{
//		c.JSON(http.StatusOK,gin.H{
//			"status":"端口必须为数字",
//			"statusCode":-1,
//		})
//		c.Abort()
//		return
//	}
//	co := conn.OraSourceConfig{Username:username,Password:password,Ip:ipaddr,Port:port,Service_name:srvname}
//	_,err = m.InitDB(co)
//	if err!=nil{
//		c.JSON(http.StatusOK,gin.H{
//			"status":"连接数据库失败",
//			"statusCode":-1,
//		})
//		c.Abort()
//		return
//	}else{
//		c.JSON(http.StatusOK,gin.H{
//			"status":"连接数据库成功",
//			"statusCode":0,
//		})
//	}
//}
//
//func AddOraConnector(c *gin.Context){
//	var ora m.OraConnectors
//
//	ora.Connector = c.PostForm("connector")
//	ora.Ipaddr = c.PostForm("ipaddr")
//	_,err := strconv.Atoi(c.PostForm("port"))
//	if err != nil{
//		c.JSON(http.StatusBadRequest,gin.H{
//			"status":"端口必须为数字",
//			"statusCode":-1,
//			"data":ora,
//		})
//		c.Abort()
//		return
//	}
//
//	ora.Port = c.PostForm("port")
//
//	ora.Srvname = c.PostForm("srvname")
//	ora.Username = c.PostForm("username")
//
//	pwdstr,err:= utils.Encrypt(c.PostForm("password"),[]byte(utils.AESKEY))
//
//	if err!=nil{
//		c.JSON(http.StatusOK,gin.H{
//			"title":"index",
//			"status":"密码加密失败",
//			"statusCode":-1,
//		})
//		c.Abort()
//		return
//	}
//
//	ora.Password = pwdstr
//
//	ora.CreateConnector()
//	c.JSON(http.StatusOK,gin.H{
//		"status":"添加成功",
//		"statusCode":0,
//		"data":ora,
//	})
//}
//
//func SaveOraConnector(c *gin.Context){
//	var ora m.OraConnectors
//
//	id,err := strconv.Atoi(c.PostForm("id"))
//	if err != nil{
//		c.JSON(http.StatusBadRequest,gin.H{
//			"status":"Id必须为数字",
//			"statusCode":-1,
//			"data":ora,
//		})
//		c.Abort()
//		return
//	}
//	ora.Id = int(id)
//	ora.Connector = c.PostForm("connector")
//	ora.Ipaddr = c.PostForm("ipaddr")
//	_,err = strconv.Atoi(c.PostForm("port"))
//	if err != nil{
//		c.JSON(http.StatusBadRequest,gin.H{
//			"status":"端口必须为数字",
//			"statusCode":-1,
//			"data":ora,
//		})
//		c.Abort()
//		return
//	}
//	ora.Port = c.PostForm("port")
//	ora.Srvname = c.PostForm("srvname")
//	ora.Username = c.PostForm("username")
//	pwdstr,err:= utils.Encrypt(c.PostForm("password"),[]byte(utils.AESKEY))
//	if err!=nil{
//		c.JSON(http.StatusOK,gin.H{
//			"title":"index",
//			"status":"密码加密失败",
//			"statusCode":-1,
//		})
//		c.Abort()
//		return
//	}
//	ora.Password = pwdstr
//
//	ora.UpdateConnector()
//
//	c.JSON(http.StatusOK,gin.H{
//		"status":"更新成功",
//		"statusCode":0,
//		"data":ora,
//	})
//
//}
//
//func DeleteOraConnector(c *gin.Context){
//	var ora m.OraConnectors
//	id ,err := strconv.Atoi(c.PostForm("id"))
//
//	if err != nil {
//		c.JSON(http.StatusBadRequest,gin.H{
//			"status":"Id必须为数字",
//			"statusCode":-1,
//			"data":ora,
//		})
//		c.Abort()
//		return
//	}
//	ora.Id = int(id)
//
//	ora.DeleteConnector()
//
//	c.JSON(http.StatusOK,gin.H{
//		"status":"删除成功",
//		"statusCode":0,
//		"data":ora,
//	})
//}
//
//func getconnect(c []conn.OraSourceConfig,id int )(conn.OraSourceConfig){
//	for _,v :=range c {
//		if (v.Id == id ){
//			return v
//		}
//	}
//	return conn.OraSourceConfig{}
//}
//
//func GetHangStat_API(c *gin.Context){
//	id ,err := strconv.Atoi(c.Query("id"))
//
//	if err != nil {
//		c.JSON(http.StatusBadRequest,gin.H{
//			"status":"Id必须为数字",
//			"statusCode":-1,
//			"data":id,
//		})
//		c.Abort()
//		return
//	}
//
//	co := getconnect(conn.OraSource,id)
//
//
//	oradb,err:=m.InitDB(co)
//
//	if err!=nil{
//		c.JSON(http.StatusOK,gin.H{
//			"status":"连接失败",
//			"statusCode":-1,
//		})
//	}else{
//		ret := m.GetHangStatus(oradb)
//		m.DestroyDB(oradb)
//
//		c.JSON(http.StatusOK,gin.H{
//			"status":"获取阻塞信息成功",
//			"data":ret,
//			"statusCode":0,
//		})
//	}
//}
//
//func GetHangStat(c *gin.Context){
//
//	id ,err := strconv.Atoi(c.Query("id"))
//
//	if err != nil {
//		c.JSON(http.StatusBadRequest,gin.H{
//			"status":"Id必须为数字",
//			"statusCode":-1,
//			"data":id,
//			"error":err,
//		})
//		c.Abort()
//		return
//	}
//
//	co := getconnect(conn.OraSource,id)
//
//	oradb,err:=m.InitDB(co)
//	if err!=nil{
//		c.JSON(http.StatusOK,gin.H{
//			"status":"连接失败",
//			"statusCode":-1,
//		})
//	}else{
//		inst_id,err:=m.GetInstanceID(oradb)
//		m.DestroyDB(oradb)
//		if err!=nil {
//			c.JSON(http.StatusOK,  gin.H{
//				"status":"获取阻塞信息失败",
//				"statusCode":-1,
//			})
//		}else{
//			c.HTML(http.StatusOK, "hangstat.html", gin.H{
//				"title":      "orahang",
//				"inst_id":    inst_id,
//				"copyright":  COPYRIGHT,
//				"statusCode": 0,
//			})
//		}
//	}
//
//}
//
//func GetSQL_API(c *gin.Context){
//	id ,err := strconv.Atoi(c.Query("id"))
//
//	if err != nil {
//		c.JSON(http.StatusBadRequest,gin.H{
//			"status":"Id必须为数字",
//			"statusCode":-1,
//			"data":id,
//		})
//		c.Abort()
//		return
//	}
//	sqlid := c.Query("sqlid")
//
//
//	co := getconnect(conn.OraSource,id)
//
//	oradb,err:=m.InitDB(co)
//	if err!=nil{
//		c.JSON(http.StatusOK,gin.H{
//			"status":"连接失败",
//			"statusCode":-1,
//		})
//	}else{
//		ret := m.GetSQL(oradb,sqlid)
//		m.DestroyDB(oradb)
//
//		c.JSON(http.StatusOK,gin.H{
//			"status":"获取SQL文本成功",
//			"data":string(ret),
//			"statusCode":0,
//		})
//	}
//}
//
//func Kill_Session_API (c *gin.Context){
//	id ,err := strconv.Atoi(c.Query("id"))
//
//	if err != nil {
//		c.JSON(http.StatusBadRequest,gin.H{
//			"status":"Id必须为数字",
//			"statusCode":-1,
//			"data":id,
//		})
//		c.Abort()
//		return
//	}
//
//	sid ,err := strconv.Atoi(c.Query("sid"))
//	if err != nil {
//		c.JSON(http.StatusBadRequest,gin.H{
//			"status":"Sid必须为数字",
//			"statusCode":-1,
//			"data":id,
//		})
//		c.Abort()
//		return
//	}
//
//	co := getconnect(conn.OraSource,id)
//
//	oradb,err:=m.InitDB(co)
//	if err!=nil {
//		c.JSON(http.StatusOK, gin.H{
//			"status":     "连接失败",
//			"statusCode": -1,
//		})
//	}else{
//		optlog,status,err := m.KillSession(oradb,sid)
//		if status {
//			logs:=m.Operation_log{time.Now().Unix(),optlog.Instid,optlog.DBNAME,optlog.IP,optlog.SQLTEXT,"success" }
//			fmt.Println(logs)
//			logs.AddLog()
//		}else{
//			logs:=m.Operation_log{time.Now().Unix(),optlog.Instid,optlog.DBNAME,optlog.IP,optlog.SQLTEXT,"failure" }
//			fmt.Println(logs)
//			logs.AddLog()
//		}
//		m.DestroyDB(oradb)
//
//		if err!=nil {
//			c.JSON(http.StatusOK,gin.H{
//				"status":"删除失败",
//				"statusCode":-1,
//			})
//		}else{
//			c.JSON(http.StatusOK,gin.H{
//				"status":"删除成功",
//				"statusCode":0,
//			})
//		}
//	}
//
//}
//
//func GetOPTLog_API (c *gin.Context){
//	var currnet_page int
//	page := c.Query("page")
//	pagesize:=c.Query("limit")
//
//	if (page==""){
//		currnet_page = 1
//	}else{
//		currnet_page,_ = strconv.Atoi(page)
//	}
//
//	pgsize,_:=strconv.Atoi(pagesize)
//
//	var cnt int
//	if cnt=m.GetLogAll();cnt >=100{
//		cnt = 100
//	}
//	l,rt:=m.GetLog(currnet_page,pgsize)
//
//	if !rt {
//		c.JSON(http.StatusOK,gin.H{
//			"status":"failed",
//			"statusCode":-1,
//		})
//	}else{
//		c.JSON(http.StatusOK,gin.H{
//			"status":"ok",
//			"data":l,
//			"statusCode":0,
//			"count":cnt,
//			"page":currnet_page,
//			"pagesize":pagesize,
//		})
//	}
///*
//	c.JSON(http.StatusOK,gin.H{
//		"status":"ok",
//		"data":conn.OraSource,
//		"count":len(conn.OraSource),
//		"page":currnet_page,
//		"pagesize":pagesize,
//		"statusCode":0,
//	})*/
//}
//
//
//func Login(c *gin.Context){
//	c.HTML(http.StatusOK,"login.html",gin.H{
//		"title":      "index",
//		"copyright":  COPYRIGHT,
//		"statusCode": 0,
//		"method":     c.Request.Method,
//		"status":     "ok",
//	})
//}
//
////login_API
//func Login_API(c *gin.Context){
//	username := c.PostForm("username")
//	password := c.PostForm("password")
//	if username == "admin" && password == "admin"{
//		log.Printf("login sucess!!")
//
//		//用户密码验证成功后，设置cookie
//		//c.SetCookie("user_cookie", string(username), 10000, "/", "localhost", false, true)
//		c.SetCookie(
//			"user_cookie", // 设置cookie的key
//			"true", // 设置cookie的值
//			6000, // 过期时间
//			"/", // 所在目录
//			c.GetHeader("Host"),  //域名
//			false,  // 是否只能通过https访问
//			false) // 是否允许别人通过js获取自己的cookie
//
//		c.JSON(http.StatusOK, gin.H{
//			"method": c.Request.Method,
//			"status": "ok",
//		})
//
//	}
//
//}
