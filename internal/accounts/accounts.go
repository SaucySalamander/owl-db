package accounts

import (
	"context"
	"database/sql"
	"log"

	"github.com/spf13/viper"
)

func open_db() {
	var conninfo = viper.GetString("db_connection")

	db, err := sql.Open("postgres", conninfo)

	if err != nil {
		panic(err)
	}

	ctx := context.TODO()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}
}
