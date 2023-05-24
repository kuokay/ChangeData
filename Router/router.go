package Router

import (
	"App/Controllers"
	"App/Middlewares"

	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {

	r.Use(Middlewares.CORSMiddleware())
	// 注册
	r.POST("/register", Controllers.Register)
	// 登录
	r.POST("/login", Controllers.Login)

	//获取url数据
	r.POST("/GetUrlData", Controllers.UrlData)

	//转换数据
	r.POST("/ChangeAllData", Controllers.ChangeAllData)
	return r
}
