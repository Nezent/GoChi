package domain

import (
	"github.com/Nezent/GoChi/errs"
	"github.com/jmoiron/sqlx"
)

type TransactionRepositoryDb struct {
	client *sqlx.DB
}

func (t TransactionRepositoryDb) MakeTransaction(transaction Transaction) (*Transaction,*errs.AppError){
	sqlInsert := "INSERT INTO transaction(account_id,transaction_type,amount,transaction_date) VALUES ($1,$2,$3,$4) RETURNING transaction_id"
	var transactionId string
	err := t.client.QueryRow(sqlInsert,transaction.AccountID,transaction.TransactionType,transaction.Amount,transaction.TransactionDate).Scan(&transactionId)
	if err != nil {
		errs.NewUnexpectedError("Unexpected Database Error")
	}
	transaction.TransactionID = transactionId
	return &transaction,nil
}

func NewTransactionRepositoryDb(dbClient *sqlx.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{client: dbClient}
}