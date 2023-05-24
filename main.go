package main

import (
	"App/Router"
	"App/common"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 获取初始化的数据库
	db := common.InitDB()
	// 延迟关闭数据库
	defer db.Close()
	// 创建路由引擎
	r := gin.Default()
	Router.CollectRoutes(r)
	// 启动服务
	panic(r.Run("0.0.0.0:8080"))
}
