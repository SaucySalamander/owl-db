package database

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func OpenDBConn() *sql.DB {
	log.Info().Msg("Connecting to db")
	var conninfo = viper.GetString("db_connection")

	db, err := sql.Open("postgres", conninfo)

	if err != nil {
		log.Fatal().Msg("Failed to connect to database")
	}

	ctx := context.TODO()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal().Msg("Failed to ping database")
	}
	log.Info().Msg("Successfully connected to db")
	return db
}
