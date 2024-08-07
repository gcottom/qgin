package middleware

import (
	"time"

	"github.com/gcottom/go-zaplog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		zaplog.InfoC(ctx, "request initiated")
		ctx.Next()
		latency := time.Since(start)
		statusCode := ctx.Writer.Status()
		zaplog.InfoC(ctx, "request completed", zap.Int("status", statusCode), zap.Duration("latency", latency))
	}
}
