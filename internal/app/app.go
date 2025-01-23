package app

import (
	"github.com/Vladislav557/auth/internal/domain/postgres"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
)

func Run() {
	defer postgres.Close()
	configInit()
	postgresInit()
	loggerInit()
}

func loggerInit() {
	var err error
	var logger *zap.Logger
	switch os.Getenv("APP_ENV") {
	case "prod":
		logger, err = zap.NewProduction()
	case "dev":
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		panic("failed to initialize zap logger: " + err.Error())
	}
	zap.ReplaceGlobals(logger)
}

func postgresInit() {
	postgres.Init(os.Getenv("DATABASE_URL"))
}

func configInit() {
	if env := os.Getenv("APP_ENV"); env == "" {
		if err := godotenv.Load(); err != nil {
			panic("Error loading .env file")
		}
	}
}
