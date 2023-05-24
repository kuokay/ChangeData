package Controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"App/Utils"
	"github.com/gin-gonic/gin"
)

type UrlInfo struct {
	Url string `json:"url" binding:"required"`
}

func UrlData(c *gin.Context) {
	var data interface{}

	var url UrlInfo
	if err := c.BindJSON(&url); err != nil {
		// 返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(url.Url,"==========")
	resp, err := http.Get(url.Url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal([]byte(body), &data); err == nil {
		c.JSON(200, gin.H{
			"message": "获取数据成功",
			"data":    data,
		})
	} else {
		c.JSON(400, gin.H{
			"message": "获取数据失败",
		})
	}

}

type ChangeData struct {
	Flags string `json:"flags" binding:"required"`
	Text  string `json:"text" binding:"required"`
}

func ChangeAllData(c *gin.Context) {
	var Data ChangeData
	if err := c.BindJSON(&Data); err != nil {
		// 返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if Data.Flags == "JSON" {
		var data []byte = []byte(Data.Text)
		result := Utils.Mainchange(data)
		

		c.JSON(200, gin.H{
			"message": "获取数据成功",
			"data":  result ,
		})
	}
}
