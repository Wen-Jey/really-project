//访问接口路由配置

package router

import (
	"admin-go-api/api/controller"
	"admin-go-api/common/config"
	"admin-go-api/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

//初始化路由

func InitRouter() *gin.Engine {
	router := gin.New()

	//宕机恢复
	router.Use(gin.Recovery())
	//跨域中间域
	router.Use(middleware.Cors())
	//图片访问路径静态文件夹可直接访问
	router.StaticFS(config.Config.ImageSettings.UploadDir, http.Dir(config.Config.ImageSettings.UploadDir))
	//log日志中间件
	router.Use(middleware.Logger())
	register(router)
	return router
}

//路由注册方法

func register(router *gin.Engine) {
	//todo 后续接口url
	router.GET("/api/captcha", controller.Captcha)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("api/login", controller.Login)
	router.POST("/api/post/add", controller.CreateSysPost)

}
