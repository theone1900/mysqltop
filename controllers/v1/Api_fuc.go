package v1

import (
	"database/sql"
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
