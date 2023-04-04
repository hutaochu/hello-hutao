package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hutaochu/hello-hutao/internal/constant"
	"github.com/hutaochu/hello-hutao/internal/utils/logger"
	"go.uber.org/zap"
	"time"
)

func RequestLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		startTime := time.Now()
		context.Next()

		method := context.Request.Method
		clientIP := context.ClientIP()
		traceId, _ := context.Get(constant.TraceId)
		url := getUrl(context)
		body := getBody(context)
		latency := time.Since(startTime)
		logger.Logger.
			With(zap.String("url", url)).
			With(zap.String("method", method)).
			With(zap.String("body", body)).
			With(zap.Duration("latency", latency)).
			With(zap.String(constant.TraceId, traceId.(string))).
			With(zap.String("client-ip", clientIP)).
			Info("request logger")
	}
}

func getUrl(c *gin.Context) string {
	path := c.Request.URL.Path
	rawQuery := c.Request.URL.RawQuery

	url := path
	if rawQuery != "" {
		url = fmt.Sprintf("%s?%s", path, rawQuery)
	}
	return url
}

func getBody(c *gin.Context) string {
	if v, ok := c.Get(gin.BodyBytesKey); ok {
		if vb, ok := v.([]byte); ok {
			return string(vb)
		}
	}
	body, _ := c.GetRawData()
	return string(body)
}
