package router

import (
	"fireurl/control"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	g1 := e.Group("/v1")
	g1.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello")
	})
	g1.GET("qq", control.Hello)
}
