package Controllers

import (
	"App/Model"
	"App/common"
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Registers struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	Email      string `json:"email" binding:"required,email"`
}

func Register(c *gin.Context) {
	var register Registers
	err := c.BindJSON(&register)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 421,
			"msg":  "参数值错误",
		})
		return
	}
	if register.Password != register.RePassword {
		c.JSON(http.StatusOK, gin.H{
			"code": 422,
			"msg":  "两次输入的密码不一致",
		})
		return
	}
	db := common.GetDB()
	Username := register.Username
	Email := register.Email
	password := register.Password

	var user Model.User

	db.Where("email =?", Email).First(&user)

	if len(user.Email) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 422,
			"msg":  "用户已存在",
		})
		return

	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// 创建用户
	newUser := Model.User{
		UserName: Username,
		Email:    Email,
		Password: string(hashedPassword),
	}
	db.Create(&newUser)

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		return
	}
	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "注册成功",
	})

}

type Logins struct {
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func Login(c *gin.Context) {

	var login Logins
	err := c.BindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数值错误",
		})
		return
	}

	db := common.GetDB()
	Email := login.Email
	password := login.Password

	var user Model.User
	db.Where("email =?", Email).First(&user)

	if len(user.Email) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 422,
			"msg":  "用户不存在",
		})
		return

	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 423,
			"msg":  "密码不正确",
		})
		return
	}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		return
	}
	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登录成功",
	})

}

