package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/hutaochu/hello-hutao/constants"
)

func SetResponse(c *gin.Context, res any) {
	c.Set(constants.ContextRes, res)
}

func GetResponse(c *gin.Context) any {
	r, _ := c.Get(constants.ContextRes)
	return r
}

func SetError(c *gin.Context, err error) {
	c.Set(constants.ContextErr, err)
}

func GetError(c *gin.Context) error {
	if err, ok := c.Get(constants.ContextErr); ok && err != nil {
		return (err).(error)
	}
	return nil
}

func SetTraceId(c *gin.Context, traceId string) {
	c.Set(constants.TraceId, traceId)
}

func GetTraceId(c *gin.Context) string {
	var traceId string
	if id, _ := c.Get(constants.TraceId); id != nil {
		traceId = (id).(string)
	}
	return traceId
}
