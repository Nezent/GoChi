package domain

import (
	"log"

	"github.com/Nezent/GoChi/errs"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) SaveAccount(account CustomerAccount) (*CustomerAccount, *errs.AppError){
	SqlInsert := "INSERT INTO account (customer_id, opening_date, account_type, amount, status) VALUES ($1,$2,$3,$4,$5) RETURNING account_id"
	var accountId string
	err := d.client.QueryRow(SqlInsert,account.CustomerID,account.OpeningDate,account.AccountType,account.Amount,account.Status).Scan(&accountId)
	if err != nil{
		log.Print(err)
		return nil, errs.NewUnexpectedError("Unexpected Database Error 3")
	}
	account.AccountID = accountId
	return &account,nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{
		client: dbClient,
	}
}