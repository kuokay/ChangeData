package Middlewares


import (
	"App/common"
	"App/Model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"

)



func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context){

		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		// 非法token
		if tokenString == "" || len(tokenString) < 7 || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}

		// 提取TOKEN部分
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}


		// 获取claims中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user Model.User
		DB.Where("id =?", userId).First(&user)
		// 将用户信息写入上下文便于读取
		c.Set("user", user)
		c.Next()
	}
	}


