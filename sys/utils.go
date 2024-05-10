package sys

import (
	"flag"

	"github.com/joho/godotenv"
)

type M map[string]string

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
