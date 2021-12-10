package routes

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"mysqltop/config"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob(config.GetEnv().TemplatePath + "/*") // html模板
	router.Use(static.Serve("/web/static", static.LocalFile("./web/static", false)))
	//if config.GetEnv().Debug {
	//	pprof.Register(router) // 性能分析工具
	//}

	//加载gin 默认logger 中间件
	router.Use(gin.Logger())


	//router.Use(middleware.Auth())


	//router.Use(handleErrors())            // 错误处理
	//router.Use(filters.RegisterSession()) // 全局session
	//router.Use(filters.RegisterCache())   // 全局cache


	//无路由和无路由函数时重定向url 到 /index
	//临时调试
	//router.NoRoute(func(c *gin.Context) {
	//	c.Redirect(http.StatusTemporaryRedirect,"/index")
	//})
	//
	//router.NoMethod(func(c *gin.Context) {
	//	c.Redirect(http.StatusTemporaryRedirect,"/index")
	//})

	RegisterApiRouter(router)

	return router
}
