package accounts

import (
	"context"
	"database/sql"

	"github.com/SaucySalamander/owl-db/pkg/proto"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
)

type Account struct {
	Id   int64
	Name string
}

var serviceName string = viper.GetString("otel.service-name")

func GetAccount(ctx context.Context, request *proto.GetAccountRequest, db_pool *sql.DB) Account {
	_, span := otel.Tracer(serviceName).Start(ctx, "DbQuery")
	defer span.End()
	log.Info().Msgf("requested id: %d", request.Id)
	result, err := db_pool.Query("SELECT account_id, account_name FROM accounts WHERE account_id=$1", request.Id)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	var account_id int64
	var account_name string
	for result.Next() {
		result.Scan(&account_id, &account_name)
	}

	log.Info().Int64("account_id", account_id).Str("account_name", account_name).Send()
	return Account{
		Id:   account_id,
		Name: account_name,
	}
}

func GetAccounts(ctx context.Context, db_pool *sql.DB) []*proto.Account {
	_, span := otel.Tracer(serviceName).Start(ctx, "DbQuery")
	defer span.End()
	result, err := db_pool.Query("SELECT account_id, account_name FROM accounts")
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	var accounts []*proto.Account
	for result.Next() {
		var account_id int64
		var account_name string
		result.Scan(&account_id, &account_name)
		accounts = append(accounts, &proto.Account{Id: account_id, Name: account_name})
	}
	return accounts
}

func CreateAccount(request *proto.CreateAccountRequest, db_pool *sql.DB) int64 {
	var id int64
	result := db_pool.QueryRow("INSERT INTO accounts (account_name) VALUES ($1) RETURNING account_id", request.Name)
	result.Scan(&id)
	return id
}

func DeleteAccount(request *proto.DeleteAccountRequest, db_pool *sql.DB) int64 {
	var id int64
	result := db_pool.QueryRow("DELETE FROM accounts where account_id=$1 RETURNING account_id", request.Id)
	result.Scan(&id)
	return id
}
