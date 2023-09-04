package accounts

import (
	"context"
	"database/sql"
	"log"

	"github.com/SaucySalamander/owl-db/pkg/proto"
	"github.com/spf13/viper"
)

type Account struct {
	Id   int64
	Name string
}

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

func GetAccount(request *proto.GetAccountRequest) Account {
	db := open_db()
	result := db.QueryRowContext(context.TODO(), "SELECT * FROM account WHERE account_id = $1", request.Id)
	var account Account
	result.Scan(&account)

	return account
}

func CreateAccount(request *proto.CreateAccountRequest) sql.Result {
	db := open_db()

	result, err := db.ExecContext(context.TODO(), "INSERT INTO account (account_name) VALUES ($1)", request.Name)

	if err != nil {
		log.Fatal(err)
	}

	return result
}

func DeleteAccount(request *proto.DeleteAccountRequest) sql.Result {
	db := open_db()

	result, err := db.ExecContext(context.TODO(), "DELETE FROM account where account_id=$1", request.Id)

	if err != nil {
		log.Fatal(err)
	}

	return result
}
