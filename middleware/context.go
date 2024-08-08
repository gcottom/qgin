package middleware

import (
	"context"

	"github.com/gcottom/go-zaplog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var InjectRequestIDCTX bool
var LogRequestID bool

func ContextMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctxt := ctx.Request.Context()
		if InjectRequestIDCTX {
			ctxt = context.WithValue(ctxt, "request_id", ctx.GetHeader(ReqIDHeader))
		}
		if InjectRequestIDCTX && LogRequestID {
			ctxt = zaplog.CreateAndInject(ctxt)
			logger := zaplog.GetLoggerFromContext(ctxt).With(zap.String("request_id", ctx.GetHeader(ReqIDHeader))).WithOptions(zap.AddCallerSkip(1))
			ctxt = context.WithValue(ctxt, "logger", logger)
		}
		ctx.Request = ctx.Request.WithContext(ctxt)
		ctx.Next()
	}
}
