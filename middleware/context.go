package middleware

import (
	"context"

	"github.com/gcottom/go-zaplog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var InjectRequestIDCTX bool
var LogRequestID bool

func ContextMiddleware(baseCtx context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if InjectRequestIDCTX {
			baseCtx = context.WithValue(baseCtx, "request_id", ctx.GetHeader("request_id"))
		}
		if InjectRequestIDCTX && LogRequestID {
			logger := zaplog.GetLoggerFromContext(baseCtx).With(zap.String("request_id", ctx.GetHeader(ReqIDHeader))).WithOptions(zap.AddCallerSkip(1))
			baseCtx = context.WithValue(baseCtx, "logger", logger)
		}
		ctx.Request = ctx.Request.WithContext(baseCtx)
		ctx.Next()
	}
}
