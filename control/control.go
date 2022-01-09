package control

import (
	"fmt"
	"log"
	"net/http"

	"fireurl/model"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, "aaaaaaa")
}

// 处理重定向请求
// 取path 到db中查询原url，再重定向到该url
func Default(c *gin.Context) {
	s := c.Param("short")
	fmt.Println(s)
	//todo chaxuan
	ourl := "www.baidu.com"
	c.Redirect(http.StatusFound, ourl)
}

// 获取短链接
func GetShortURL(c *gin.Context) {
	d := new(model.URLData)
	ourl, ok := c.GetPostForm("origin_url")
	if !ok {
		log.Fatal("no param ourl")
	}
	time, ok := c.GetPostForm("exp_time")
	if !ok {
		log.Fatal("no param exptime")
	}
	d.OriginURL = ourl
	fmt.Println(ourl)
	fmt.Println(time)
	// 长化短
	//	shortURL := lib.Trans2short(ourl)

	// model.SaveURL()
}
