package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewServer(port int) {
	g := gin.Default()
	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "welcome to my api"})
	})
	g.Run(fmt.Sprintf(":%v", port))
}
