package initialize

import (
	"snowy-video-serve/global"
	"snowy-video-serve/middleware"
	"snowy-video-serve/router"

	//"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	//Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.SYS_LOG.Info("use middleware logger")

	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.SYS_LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.SYS_LOG.Info("register swagger handler")

	Router.Static("/videosPath", "G:\\MyDrivers\\hotfix\\国\\新建文件夹\\")
	// 方便统一添加路由组前缀 多服务器上线使用

	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}

	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		router.InitUserRouter(PrivateGroup)  // 注册用户路由
		router.InitVideoRouter(PrivateGroup) // 注册视频路由
	}
	global.SYS_LOG.Info("router register success")

	return Router
}
