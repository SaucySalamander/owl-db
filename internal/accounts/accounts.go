package accounts

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/SaucySalamander/owl-db/pkg/proto"
)

type Account struct {
	Id   int64
	Name string
}

func GetAccount(request *proto.GetAccountRequest, db_pool *sql.DB) Account {
	fmt.Println("requested id:", request.Id)
	result, err := db_pool.Query("SELECT account_id, account_name FROM account WHERE account_id=$1", request.Id)
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

func CreateAccount(request *proto.CreateAccountRequest, db_pool *sql.DB) int64 {
	var id int64
	result := db_pool.QueryRow("INSERT INTO account (account_name) VALUES ($1) RETURNING account_id", request.Name)
	result.Scan(&id)
	return id
}

func DeleteAccount(request *proto.DeleteAccountRequest, db_pool *sql.DB) int64 {
	var id int64
	result := db_pool.QueryRow("DELETE FROM account where account_id=$1 RETURNING account_id", request.Id)
	result.Scan(&id)
	return id
}
