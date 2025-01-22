package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var DB *sql.DB

func Init(url string) {
	zap.L().Info("starting to connect to postgres")
	var err error
	DB, err = sql.Open("postgres", url)
	if err != nil {
		panic("failed to connect database " + err.Error())
	}
	if err = DB.Ping(); err != nil {
		panic("failed to ping database " + err.Error())
	}
	zap.L().Info("successfully connected to postgres")
}

func Close() {
	zap.L().Info("closing database connection")
	if err := DB.Close(); err != nil {
		panic("failed to close database " + err.Error())
	}
	zap.L().Info("successfully closed database connection")
}
