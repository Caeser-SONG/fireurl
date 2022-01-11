package control

import (
	"fmt"
	"log"
	"net/http"

	"fireurl/lib"
	"fireurl/model"
	"strconv"

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
	//todo chaxun
	//ourl := "http://81.70.91.175:8888/v1/test"
	ourl := "https://baidu.com"
	c.Redirect(http.StatusMovedPermanently, ourl)
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
		fmt.Println("no param exptime")
	}

	d.OriginURL = ourl
	time1, err := strconv.ParseInt(time, 10, 64)
	if err != nil {
		fmt.Println("time tans into int64 fail")
	}
	d.Exptime = time1
	fmt.Println(ourl)
	fmt.Println(time)

	// 长化短
	shortURL := lib.Trans2short(ourl)
	d.ShortURL = shortURL
	model.SaveURL(d)
	c.JSON(http.StatusOK, d)

}
