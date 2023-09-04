package accounts

import (
	"context"
	"database/sql"
	"log"

	"github.com/SaucySalamander/owl-db/pkg/proto/account"
	"github.com/spf13/viper"
)

func open_db() *sql.DB {
	var conninfo = viper.GetString("db_connection")

	db, err := sql.Open("postgres", conninfo)

	if err != nil {
		panic(err)
	}

	ctx := context.TODO()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	return db
}

func CreateAccount(request *account.CreateAccountRequest) sql.Result {
	db := open_db()

	result, err := db.ExecContext(context.TODO(), "INSERT INTO account (name) VALUES ($1)", request.Name)

	if err != nil {
		log.Fatal(err)
	}

	return result
}
