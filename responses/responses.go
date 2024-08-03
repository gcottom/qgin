package responses

import (
	"github.com/gcottom/qgin/qgin"
	"github.com/gin-gonic/gin"
)

type Failure struct {
	RequestID string `json:"request_id,omitempty"`
	Message   string `json:"message,omitempty"`
}

type Success struct {
	RequestID string `json:"request_id,omitempty"`
	Data      any    `json:"data,omitempty"`
}

func FailureResponse(ctx *gin.Context, statusCode int, message string) {
	if qgin.GetActiveConfig().UseRequestIDMW {
		ctx.JSON(statusCode, Failure{
			RequestID: ctx.GetString("request_id"),
			Message:   message,
		})
		return
	}
	ctx.JSON(statusCode, Failure{
		Message: message,
	})
}

func SuccessResponse(ctx *gin.Context, statusCode int, data any) {
	if qgin.GetActiveConfig().UseRequestIDMW {
		ctx.JSON(statusCode, Success{
			RequestID: ctx.GetString("request_id"),
			Data:      data,
		})
		return
	}
	ctx.JSON(statusCode, Success{
		Data: data,
	})
}

func SuccessResponseNoContent(ctx *gin.Context) {
	ctx.JSON(204, nil)
}
