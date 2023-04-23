package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hutaochu/hello-hutao/constants"
	"github.com/hutaochu/hello-hutao/pkg/utils"
)

func Trace() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestId := context.GetHeader(constants.RequestIdHeader)
		if requestId == "" {
			utils.SetTraceId(context, uuid.New().String())
		} else {
			utils.SetTraceId(context, requestId)
		}

		context.Next()
	}
}
