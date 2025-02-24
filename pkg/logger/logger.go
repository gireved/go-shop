package logger

import (
	"go.uber.org/zap"
)

var loggerInstance *zap.Logger

// 初始化日志记录器
func InitLog() error {
	var err error
	loggerInstance, err = zap.NewProduction() // 生产环境日志格式
	if err != nil {
		return err
	}
	return nil
}

// Sync 关闭日志记录器
func Sync() {
	if loggerInstance != nil {
		_ = loggerInstance.Sync()
	}
}

func Info(message string, fields ...zap.Field) {
	if loggerInstance != nil {
		loggerInstance.Info(message, fields...)
	}
}

func Debug(message string, fields ...zap.Field) {
	if loggerInstance != nil {
		loggerInstance.Debug(message, fields...)
	}
}

func Error(message string, fields ...zap.Field) {
	if loggerInstance != nil {
		loggerInstance.Error(message, fields...)
	}
}

func Fatal(message string, fields ...zap.Field) {
	if loggerInstance != nil {
		loggerInstance.Fatal(message, fields...)
	}
}
