package router

import (
	"fireurl/control"
	"fireurl/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	//e.Use(middleware.Check())
	g1 := e.Group("/v1")
	g1.Use(middleware.Check())
	g1.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello")
	})
	g1.GET("test", control.Hello)
	g1.POST("create", control.GetShortURL)

	e.GET("/:short", control.Default)
}
