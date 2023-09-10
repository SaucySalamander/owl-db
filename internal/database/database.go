package database

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func OpenDBConn() *sql.DB {
	var conninfo = viper.GetString("db_connection")

	db, err := sql.Open("postgres", conninfo)

	if err != nil {
		panic(err)
	}

	ctx := context.TODO()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal().Msg("Failed to ping database")
	}

	return db
}
