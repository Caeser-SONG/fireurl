package main

import (
	"fmt"

	"fireurl/router"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("vim-go")
	r := gin.Default()
	router.RegisterRouter(r)
	r.Run(":8888")
}
