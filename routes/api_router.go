package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mysqltop/controllers/v1"

	//"mysqltop/middleware"
	"net/http"
	"strings"
)

//先设计html 页面，
//设置api_router.go 中路由，
//设置Api_fuc.go 中渲染函数

//api/init.json 配置主页面左侧向导url

func RegisterApiRouter(router *gin.Engine) {
	//使用CORS （Cross-Origin Resource Sharing，跨域资源共享）
	router.Use(Cors())

	//router.GET("/login", v1.Login)
	//router.POST("/login_API",v1.Login_API)

	//页面接口
	router.GET("/index", v1.Index)
	router.GET("/dblist", v1.Dblist)
	router.GET("/dblistadd", v1.Dblistadd)
	router.GET("/dblistedit", v1.Dblistedit)
	router.GET("/topprocesslist", v1.Topprocesslist)
	router.GET("/healcheck", v1.Healthcheck)
	router.GET("/seccheck", v1.Seccheck)
	router.GET("/returnhcreport", v1.Returnhcreport)
	router.GET("/mysqltopsess", v1.Mysqltopsess)
	//slow sql
	router.GET("/mysqlslowsqllist", v1.Mysqlslowsqllist)
	router.GET("/mysqlslowsql", v1.Mysqlslowsql)

	router.GET("/dbinfo", v1.Dbinfo1)
	//router.GET("/getdbinfo",v1.Getdbinfo)

	router.GET("/about", v1.About)

	//数据接口，注意接口请求方式：post or get
	router.GET("/GetUsers", v1.GetUsers)
	router.POST("/Getmysqltopsess", v1.Getmysqltopsess)
	router.POST("/Getmysqltopsql", v1.Getmysqltopsql)

	//router.GET("/getdblist",v1.Getdblist)
	router.POST("/postadddblist", v1.Postadddblist)
	router.POST("/postupdatedblist", v1.Postupdatedblist)
	router.POST("/postdblistdel", v1.Postdeldblist)
	router.POST("/postfinddblist", v1.Finddblist)

	//router.GET("/hahahahaha", middleware.Auth(),v1.HAHAHAHA)
	//router.GET("/getora",middleware.Auth(),v1.ListOraConnector)
	//router.GET("/getora2",middleware.Auth(),v1.ListOraConnector2)
	//router.POST("/addora",middleware.Auth(),v1.AddOraConnector)
	//router.GET("/addPage",middleware.Auth(),v1.AddOraConnPage)
	//router.POST("/saveora",middleware.Auth(),v1.SaveOraConnector)
	//router.GET("/savePage",middleware.Auth(),v1.SaveOraConnPage)
	//router.POST("/delora",middleware.Auth(),v1.DeleteOraConnector)
	//router.POST("/testora",middleware.Auth(),v1.TryConnectin)
	//router.GET("/listhang",middleware.Auth(),v1.GetHangStat)
	//router.GET("/gethang",middleware.Auth(),v1.GetHangStat_API)
	//router.GET("/gethangs",middleware.Auth(),v1.GetHangs)
	//router.GET("/getsql",middleware.Auth(),v1.GetSQL_API)
	//router.GET("/killsess",middleware.Auth(),v1.Kill_Session_API)
	//router.GET("/loginfo",middleware.Auth(),v1.LoginfoPage)
	//router.GET("/logs",middleware.Auth(),v1.GetOPTLog_API)
}

//CORS （Cross-Origin Resource Sharing，跨域资源共享）是一个系统，它由一系列传输的HTTP头组成，这些HTTP头决定浏览器是否阻止前端 JavaScript 代码获取跨域请求的响应。

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		c.Header("Access-Control-Allow-Origin", "*")
		//c.Header("Access-Control-Allow-Headers", headerStr)
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "X-ACCESS_TOKEN, Authorization, Content-Length, X-CSRF-Token, Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		// c.Header("Access-Control-Max-Age", "172800")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Set("content-type", "application/json")
		//c.Set("X-Frame-Options", "SAMEORIGIN")
		//c.Writer.Header().Set("X-Frame-Options", "deny")
		if origin != "" {
			// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.AbortWithStatus(http.StatusNoContent)
			//c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next()
	}
}
