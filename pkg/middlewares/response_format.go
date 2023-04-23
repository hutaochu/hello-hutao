package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hutaochu/hello-hutao/pkg/utils"
	"github.com/hutaochu/hello-hutao/pkg/utils/logger"
	"net/http"
)

type successResponse struct {
	Version string
	Success bool
	Result  any
}

type errorResponse struct {
	Version    string
	Success    bool
	ErrMessage string
	Error      error
}

func HandleResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if r := utils.GetResponse(c); r != nil {
			c.JSON(http.StatusOK, successResponse{
				Version: "1.0.0",
				Success: true,
				Result:  r,
			})
			return
		}
		if e := utils.GetError(c); e != nil {
			err := (e).(error)
			logger.Error(c, err, "handle response get error")
			c.JSON(http.StatusInternalServerError, errorResponse{
				Version:    "1.0.0",
				Success:    false,
				Error:      err,
				ErrMessage: err.Error(),
			})
			return
		}

		err := errors.New("not found")
		c.JSON(http.StatusNotFound, errorResponse{
			Version:    "1.0.0",
			Success:    false,
			Error:      err,
			ErrMessage: err.Error(),
		})
	}
}
