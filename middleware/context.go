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
		var logger *zap.Logger
		if InjectRequestIDCTX && LogRequestID {
			logger = zaplog.GetLoggerFromContext(baseCtx).With(zap.String("request_id", ctx.GetHeader(ReqIDHeader))).WithOptions(zap.AddCallerSkip(1))
		}
		ctx.Request = ctx.Request.WithContext(context.WithValue(context.WithValue(ctx.Request.Context(), "logger", logger), "request_id", ctx.GetHeader(ReqIDHeader)))
		ctx.Next()
	}
}
