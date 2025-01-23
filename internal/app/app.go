package app

import (
	"context"
	"github.com/Vladislav557/auth/internal/resources"
	"github.com/Vladislav557/auth/internal/resources/postgres"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
)

func Run() {
	defer postgres.Close()
	configInit()
	postgresInit()
	loggerInit()
	serverInit()
}

func serverInit() {
	zap.L().Info("starting server")
	r := resources.RouterInit()
	s := resources.New("8080", r)
	go func() {
		err := s.Start()
		if err != nil {
			panic("failed to start server")
		}
	}()
	quitSigCh := make(chan os.Signal, 1)
	sig := <-quitSigCh
	zap.S().Info("Service stopped by signal ", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown: " + err.Error())
	}
}

func loggerInit() {
	var err error
	var logger *zap.Logger
	if appEnv := os.Getenv("APP_ENV"); appEnv == "" {
		panic("env variable APP_ENV is not set")
	}
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
