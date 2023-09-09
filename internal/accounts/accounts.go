package accounts

import (
	"context"
	"database/sql"
	"fmt"
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
	fmt.Println("requested id:", request.Id)
	result, err := db.Query("SELECT account_id, account_name FROM account WHERE account_id=$1", request.Id)
	fmt.Println(result.Columns())
	if err != nil {
		log.Fatal(err)
	}
	var account_id int64
	var account_name string
	for result.Next() {
		result.Scan(&account_id, &account_name)
	}

	fmt.Println(account_id)
	fmt.Println(account_name)
	return Account{
		Id:   account_id,
		Name: account_name,
	}
}

func CreateAccount(request *proto.CreateAccountRequest) int64 {
	db := open_db()
	var id int64
	result := db.QueryRow("INSERT INTO account (account_name) VALUES ($1) RETURNING account_id", request.Name)
	result.Scan(&id)
	return id
}

func DeleteAccount(request *proto.DeleteAccountRequest) int64 {
	db := open_db()
	var id int64
	result := db.QueryRow("DELETE FROM account where account_id=$1 RETURNING account_id", request.Id)
	result.Scan(&id)
	return id
}
