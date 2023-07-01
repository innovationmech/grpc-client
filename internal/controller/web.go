package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/innovationmech/grpc-client/internal/pool"
)

var r = gin.Default()

func Serve() {
	r.Run(":8080")
}

func init() {
	r.GET("/health", health)
	r.GET("/dial_hello", dial)
}
