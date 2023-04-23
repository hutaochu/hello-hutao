package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hutaochu/hello-hutao/constants"
	"github.com/hutaochu/hello-hutao/pkg/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger
var logger *zap.Logger

func init() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/info.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
		Compress:   true,
	})
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.InfoLevel,
	)
	logger = zap.New(core)
	Logger = logger
}

func Debug(c *gin.Context, msg string, fields ...zap.Field) {
	logger.With(zap.String(constants.TraceId, utils.GetTraceId(c))).Debug(msg, fields...)
}

func Info(c *gin.Context, msg string, fields ...zap.Field) {
	logger.With(zap.String(constants.TraceId, utils.GetTraceId(c))).Info(msg, fields...)
}

func Warn(c *gin.Context, msg string, fields ...zap.Field) {
	logger.With(zap.String(constants.TraceId, utils.GetTraceId(c))).Warn(msg, fields...)
}

func Error(c *gin.Context, err error, msg string, fields ...zap.Field) {
	logger.
		With(zap.String("error", fmt.Sprint(err))).
		With(zap.String(constants.TraceId, utils.GetTraceId(c))).
		Error(msg, fields...)
}

func Fatal(c *gin.Context, msg string, fields ...zap.Field) {
	logger.With(zap.String(constants.TraceId, utils.GetTraceId(c))).Fatal(msg, fields...)
}
