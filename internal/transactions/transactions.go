package transactions

import (
	"database/sql"

	"github.com/SaucySalamander/owl-db/pkg/proto"
	"github.com/rs/zerolog/log"
)

func GetTransaction(request *proto.GetTransactionRequest, db_pool *sql.DB) *proto.Transaction {
	var transaction proto.Transaction
	result, err := db_pool.Query("SELECT transaction_id, transaction_timestamp, account_id, category_id, subcategory_id, transaction_name, credit, debit FROM transactions WHERE transaction_id=$1", request.Id)
	if err != nil {
		log.Warn().Err(err).Send()
	}

	for result.Next() {
		err = result.Scan(&transaction.Id, &transaction.Timestamp.Seconds, &transaction.AccountId, &transaction.CategoryId, &transaction.SubCategoryId, &transaction.Name, &transaction.Credit, &transaction.Debit)
		if err != nil {
			log.Warn().Err(err).Send()
		}
	}
	return &transaction
}

func CreateTransaction(request *proto.CreateTransactionRequest, db_pool *sql.DB) int64 {
	var id int64
	log.Info().Msgf("%v", request)
	result := db_pool.QueryRow("INSERT INTO transactions (transaction_timestamp,account_id,category_id,subcategory_id,transaction_name,credit,debit) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING transaction_id", request.Transaction.Timestamp.AsTime().Unix(), request.Transaction.AccountId, request.Transaction.CategoryId, request.Transaction.SubCategoryId, request.Transaction.Name, request.Transaction.Credit, request.Transaction.Debit)
	err := result.Scan(&id)
	if err != nil {
		log.Warn().Err(err).Send()
	}
	return id
}

func DeleteTransaction(request *proto.DeleteTransactionRequest, db_pool *sql.DB) int64 {
	var id int64
	result := db_pool.QueryRow("DELETE FROM transactions where transaction_id=$1 RETURNING transaction_id", request.Id)
	result.Scan(&id)
	return id
}
