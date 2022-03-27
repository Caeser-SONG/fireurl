package main

import (
	"fireurl/model"
	"fireurl/router"
	"fmt"

	"github.com/gin-contrib/pprof"

	"github.com/gin-gonic/gin"
)

func main() {

	model.InitMongo()
	fmt.Println("init mongo success")
	r := gin.Default()
	pprof.Register(r)
	router.RegisterRouter(r)
	fmt.Println("registe success")
	r.Run(":8888")
}
