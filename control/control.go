package control

import (
	"fireurl/lib"
	"fireurl/model"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, "hello world")
}

// 处理重定向请求
// 取path 到db中查询原url，再重定向到该url
func Default(c *gin.Context) {
	s := c.Param("short")

	ourl, err := model.GetOriginURL(s)
	if err != nil {
		c.JSON(http.StatusOK, err)
	}
	//todo chaxun
	c.Redirect(http.StatusFound, ourl)
	//c.JSON(200, ourl)
}

// 生成短链接
func GetShortURL(c *gin.Context) {
	d := new(model.URLData)
	ourl, ok := c.GetPostForm("origin_url")
	if !ok {
		log.Fatal("no param ourl")
	}
	time, ok := c.GetPostForm("exp_time")
	if !ok {
		fmt.Println("no param exptime")
	}
	d.OriginURL = ourl

	time1, err := strconv.ParseInt(time, 10, 64)
	if err != nil {
		fmt.Println("time tans into int64 fail")
	}
	d.Exptime = time1
	var transer lib.Transer
	transer = new(lib.TransMD5)
	// 长变短
	shortURL := transer.Trans2Short(ourl)
	// find 是否存在
	d.ShortURL = shortURL
	if model.IsExist(shortURL) {
		c.JSON(http.StatusOK, d)
		return
	}
	model.SaveURL(d)
	c.JSON(http.StatusOK, d)
}
