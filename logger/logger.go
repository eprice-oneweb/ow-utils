package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/eprice-oneweb/ow-utils/envutil"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger = generateLogger()
var LOG_LEVEL = envutil.GetEnvOrDefault("LOG_LEVEL", "DEBUG")

func generateLogger() *zap.SugaredLogger {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(config)

	logLevel, err := zapcore.ParseLevel(LOG_LEVEL)
	if err != nil {
		var errMsg = "error parsing LOG_LEVEL environment variable of: " + LOG_LEVEL + " with error: " + err.Error()
		log.Fatalln(errMsg)
	}
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logLevel),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return logger.Sugar()
}

// logFormatter is the log format function for this project for gin messages
var LogFormatterFunc = func(param gin.LogFormatterParams) string {
	Logger.Info(fmt.Sprintf("%s %v completed with status: %d latency: %v client IP: %s",
		param.Method,
		param.Path,
		param.StatusCode,
		param.Latency,
		param.ClientIP,
	))
	return ""
}
