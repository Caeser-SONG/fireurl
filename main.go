package main

import (
	"fireurl/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.RegisterRouter(r)
	r.Run(":8888")
}
