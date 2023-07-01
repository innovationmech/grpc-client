package controller

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/innovationmech/grpc-client/internal/pool"
	"github.com/innovationmech/grpc-client/pb"
)

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
