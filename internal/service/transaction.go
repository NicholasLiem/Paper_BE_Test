package service

import (
	"github.com/NicholasLiem/Paper_BE_Test/internal/datastruct"
	"github.com/NicholasLiem/Paper_BE_Test/internal/repository"
	"github.com/NicholasLiem/Paper_BE_Test/utils"
	"log"
	"net/http"
)

type TransactionService interface {
	CreateTransaction(transaction datastruct.Transaction) (bool, *utils.HttpError)
	GetTransaction(transactionID uint) (*datastruct.Transaction, *utils.HttpError)
	GetAllTransactions() ([]datastruct.Transaction, *utils.HttpError)
	GetTransactionsByWalletID(walletID uint) ([]datastruct.Transaction, *utils.HttpError)
}

type transactionService struct {
	dao repository.DAO
}

func NewTransactionService(dao repository.DAO) TransactionService {
	return &transactionService{dao: dao}
}

func (ts *transactionService) CreateTransaction(transaction datastruct.Transaction) (bool, *utils.HttpError) {
	if transaction.Amount <= 0 {
		return false, &utils.HttpError{Message: "Invalid transaction amount", StatusCode: http.StatusBadRequest}
	}

	success, err := ts.dao.NewTransactionQuery().CreateTransaction(transaction)
	if err != nil {
		log.Printf("Error creating transaction: %v", err)
		return false, &utils.HttpError{Message: "Error creating transaction", StatusCode: http.StatusInternalServerError}
	}
	return success, nil
}

func (ts *transactionService) GetTransaction(transactionID uint) (*datastruct.Transaction, *utils.HttpError) {
	transaction, err := ts.dao.NewTransactionQuery().GetTransaction(transactionID)
	if err != nil {
		log.Printf("Error retrieving transaction: %v", err)
		return nil, &utils.HttpError{Message: "Error retrieving transaction", StatusCode: http.StatusInternalServerError}
	}
	return transaction, nil
}

func (ts *transactionService) GetAllTransactions() ([]datastruct.Transaction, *utils.HttpError) {
	transactions, err := ts.dao.NewTransactionQuery().GetAllTransactions()
	if err != nil {
		log.Printf("Error retrieving transactions: %v", err)
		return nil, &utils.HttpError{Message: "Error retrieving transactions", StatusCode: http.StatusInternalServerError}
	}
	return transactions, nil
}

func (ts *transactionService) GetTransactionsByWalletID(walletID uint) ([]datastruct.Transaction, *utils.HttpError) {
	transactions, err := ts.dao.NewTransactionQuery().GetTransactionsByWalletID(walletID)
	if err != nil {
		log.Printf("Error retrieving transactions: %v", err)
		return nil, &utils.HttpError{Message: "Error retrieving transactions", StatusCode: http.StatusInternalServerError}
	}
	return transactions, nil
}
