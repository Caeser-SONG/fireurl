package main

import (
	"fireurl/model"
	"fireurl/router"

	"github.com/gin-gonic/gin"
)

func main() {
	model.InitMongo()
	r := gin.Default()
	router.RegisterRouter(r)
	r.Run(":8888")
}
