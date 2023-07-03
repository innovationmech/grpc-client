package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var limitChan = make(chan struct{}, 200)

func TimerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		begin := time.Now()
		c.Next()
		fmt.Printf("Request took %v\n", time.Since(begin).Milliseconds())
	}
}

func LimitMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limitChan <- struct{}{}
		ctx.Next()
		<-limitChan
	}
}
