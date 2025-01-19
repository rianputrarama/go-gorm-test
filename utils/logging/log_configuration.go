package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	serviceLog = "logs/api_service.log"
	httpLog    = "logs/http_request.log"
)

var (
	Logger     *zap.Logger
	HttpLogger *zap.Logger
)

func Init() {
	serviceWriteSyncer := getLogWriter(serviceLog)
	httpsWriteSyncer := getLogWriter(httpLog)

	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	core := zapcore.NewCore(encoder, serviceWriteSyncer, zapcore.DebugLevel)
	httpCore := zapcore.NewCore(encoder, httpsWriteSyncer, zapcore.DebugLevel)

	Logger = zap.New(core)
	HttpLogger = zap.New(httpCore)
}

func getLogWriter(logFile string) zapcore.WriteSyncer {
	file, _ := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	return zapcore.AddSync(file)
}
