package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const ReqIDHeader = "X-Request-ID"

func UUIDGenerator() string {
	return uuid.New().String()
}

func RequestIDMiddleware(generator func() string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := ctx.GetHeader(ReqIDHeader)
		if requestID == "" {
			requestID = generator()
		}
		ctx.Request.Header.Set(ReqIDHeader, requestID)
		ctx.Header(ReqIDHeader, requestID)
		ctx.Set("request_id", requestID)
		ctx.Next()
	}
}
