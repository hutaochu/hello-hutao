package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hutaochu/hello-hutao/internal/constant"
)

func Trace() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestId := context.GetHeader(constant.RequestIdHeader)
		if requestId == "" {
			context.Set(constant.TraceId, uuid.New().String())
		} else {
			context.Set(constant.TraceId, requestId)
		}

		context.Next()
	}
}
