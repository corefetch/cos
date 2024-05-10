package sys

import (
	"flag"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type M map[string]any

func Logger() *zap.SugaredLogger {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	return logger.Sugar()
}

func LoadEnv() {

	devMode := flag.Bool("dev", false, "Run in development mode")

	flag.Parse()

	var err error

	if *devMode {
		err = godotenv.Load(".env.dev")
	} else {
		err = godotenv.Load()
	}

	if err != nil {
		panic(err)
	}
}
