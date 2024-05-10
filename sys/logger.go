package sys

import "go.uber.org/zap"

func Logger() *zap.SugaredLogger {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	return logger.Sugar()
}
