package controller

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/innovationmech/grpc-client/internal/middleware"
	"github.com/innovationmech/grpc-client/internal/pool"
	"github.com/innovationmech/grpc-client/pb"
	"github.com/spf13/cobra"
)

var r = gin.Default()

func NewServeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			r.Use(middleware.LimitMiddleware())
			r.Use(middleware.TimerMiddleware())
			r.GET("/health", health)
			r.GET("/dial_hello", dial)
			r.Run(":8080")
		},
	}
}

func health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func dial(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": hello(),
	})
}

func hello() string {
	cp := pool.GetDefaultConnPool()
	conn := cp.Get()
	if conn == nil {
		return ""
	}
	c := pb.NewHelloServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Greeting: "world"})
	if err != nil {
		return err.Error()
	}
	return r.Reply
}
