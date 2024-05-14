package repository

import (
	"errors"
	"github.com/NicholasLiem/Paper_BE_Test/internal/datastruct"
	"gorm.io/gorm"
)

type TransactionQuery interface {
	BeginTransaction() *gorm.DB
	CreateTransaction(transaction datastruct.Transaction) (bool, error)
	CreateTransactionTx(transaction datastruct.Transaction, tx *gorm.DB) (bool, error)
	GetTransaction(transactionID uint) (*datastruct.Transaction, error)
	GetAllTransactions() ([]datastruct.Transaction, error)
	GetTransactionsByWalletID(walletID uint) ([]datastruct.Transaction, error)
}

type transactionQuery struct {
	pgdb *gorm.DB
}

func NewTransactionQuery(db *gorm.DB) TransactionQuery {
	return &transactionQuery{pgdb: db}
}

func (tq *transactionQuery) BeginTransaction() *gorm.DB {
	return tq.pgdb.Begin()
}

func (tq *transactionQuery) CreateTransaction(transaction datastruct.Transaction) (bool, error) {
	if err := validateTransactionType(transaction.Type); err != nil {
		return false, err
	}

	result := tq.pgdb.Create(&transaction)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (tq *transactionQuery) CreateTransactionTx(transaction datastruct.Transaction, tx *gorm.DB) (bool, error) {
	if err := validateTransactionType(transaction.Type); err != nil {
		return false, err
	}

	result := tx.Create(&transaction)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (tq *transactionQuery) GetTransaction(transactionID uint) (*datastruct.Transaction, error) {
	var transaction datastruct.Transaction
	result := tq.pgdb.First(&transaction, transactionID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &transaction, nil
}

func (tq *transactionQuery) GetAllTransactions() ([]datastruct.Transaction, error) {
	var transactions []datastruct.Transaction
	result := tq.pgdb.Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}

func (tq *transactionQuery) GetTransactionsByWalletID(walletID uint) ([]datastruct.Transaction, error) {
	var transactions []datastruct.Transaction
	result := tq.pgdb.Where("wallet_id = ?", walletID).Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}

func validateTransactionType(transactionType datastruct.TransactionType) error {
	if transactionType != datastruct.CashIn && transactionType != datastruct.CashOut {
		return errors.New("invalid transaction type")
	}
	return nil
}
