package middleware

import (
	"fireurl/constant"

	"github.com/gin-gonic/gin"
)

func Check() gin.HandlerFunc {
	return func(c *gin.Context) {

		sign := c.Query("sign")
		if sign != constant.Key {
			c.Abort()
			c.JSON(200, "sign error")
		} else {
			c.Next()
		}

	}
}
