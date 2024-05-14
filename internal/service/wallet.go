package service

import (
	"github.com/NicholasLiem/Paper_BE_Test/internal/datastruct"
	"github.com/NicholasLiem/Paper_BE_Test/internal/repository"
	"github.com/NicholasLiem/Paper_BE_Test/utils"
	"log"
	"net/http"
	"sync"
)

type WalletService interface {
	Topup(userID uint, amount float64) (bool, *utils.HttpError)
	Withdraw(userID uint, amount float64) (bool, *utils.HttpError)
	GetBalance(userID uint) (float64, *utils.HttpError)
	GetTransactions(userID uint) ([]datastruct.Transaction, *utils.HttpError)
}

type walletService struct {
	dao                repository.DAO
	transactionService TransactionService
	mu                 sync.Mutex
}

func NewWalletService(dao repository.DAO, transactionService TransactionService) WalletService {
	return &walletService{dao: dao, transactionService: transactionService}
}

func (ws *walletService) Topup(userID uint, amount float64) (bool, *utils.HttpError) {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	if amount <= 0 {
		return false, &utils.HttpError{Message: "Invalid topup amount", StatusCode: http.StatusBadRequest}
	}

	tx := ws.dao.NewWalletQuery().BeginTransaction()
	if tx.Error != nil {
		return false, &utils.HttpError{Message: "Error starting database transaction", StatusCode: http.StatusInternalServerError}
	}

	wallet, err := ws.dao.NewWalletQuery().GetWalletByUserID(userID)
	if err != nil {
		tx.Rollback()
		log.Printf("Error retrieving wallet: %v", err)
		return false, &utils.HttpError{Message: "Error retrieving wallet", StatusCode: http.StatusInternalServerError}
	}

	if wallet == nil {
		wallet = &datastruct.Wallet{UserID: userID, Balance: amount}
		if success, err := ws.dao.NewWalletQuery().CreateWalletTx(*wallet, tx); !success || err != nil {
			tx.Rollback()
			log.Printf("Error creating wallet: %v", err)
			return false, &utils.HttpError{Message: "Error creating wallet", StatusCode: http.StatusInternalServerError}
		}
	} else {
		wallet.Balance += amount
		if success, err := ws.dao.NewWalletQuery().UpdateWalletTx(wallet.UserID, *wallet, tx); !success || err != nil {
			tx.Rollback()
			log.Printf("Error updating wallet balance: %v", err)
			return false, &utils.HttpError{Message: "Error updating wallet balance", StatusCode: http.StatusInternalServerError}
		}
	}

	transaction := datastruct.Transaction{
		WalletID: wallet.UserID,
		Amount:   amount,
		Type:     datastruct.CashIn,
	}
	if success, err := ws.transactionService.CreateTransaction(transaction); !success || err != nil {
		tx.Rollback()
		log.Printf("Error creating transaction: %v", err)
		return false, &utils.HttpError{Message: "Error creating transaction", StatusCode: http.StatusInternalServerError}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		log.Printf("Error committing transaction: %v", err)
		return false, &utils.HttpError{Message: "Error committing transaction", StatusCode: http.StatusInternalServerError}
	}

	return true, nil
}

func (ws *walletService) Withdraw(userID uint, amount float64) (bool, *utils.HttpError) {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	if amount <= 0 {
		return false, &utils.HttpError{Message: "Invalid withdrawal amount", StatusCode: http.StatusBadRequest}
	}

	tx := ws.dao.NewWalletQuery().BeginTransaction()
	if tx.Error != nil {
		return false, &utils.HttpError{Message: "Error starting database transaction", StatusCode: http.StatusInternalServerError}
	}

	wallet, err := ws.dao.NewWalletQuery().GetWalletByUserID(userID)
	if err != nil {
		tx.Rollback()
		log.Printf("Error retrieving wallet: %v", err)
		return false, &utils.HttpError{Message: "Error retrieving wallet", StatusCode: http.StatusInternalServerError}
	}

	if wallet == nil || wallet.Balance < amount {
		tx.Rollback()
		return false, &utils.HttpError{Message: "Insufficient balance", StatusCode: http.StatusBadRequest}
	}

	wallet.Balance -= amount
	if success, err := ws.dao.NewWalletQuery().UpdateWalletTx(wallet.UserID, *wallet, tx); !success || err != nil {
		tx.Rollback()
		log.Printf("Error updating wallet balance: %v", err)
		return false, &utils.HttpError{Message: "Error updating wallet balance", StatusCode: http.StatusInternalServerError}
	}

	transaction := datastruct.Transaction{
		WalletID: wallet.UserID,
		Amount:   amount,
		Type:     datastruct.CashOut,
	}
	if success, err := ws.transactionService.CreateTransaction(transaction); !success || err != nil {
		tx.Rollback()
		log.Printf("Error creating transaction: %v", err)
		return false, &utils.HttpError{Message: "Error creating transaction", StatusCode: http.StatusInternalServerError}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		log.Printf("Error committing transaction: %v", err)
		return false, &utils.HttpError{Message: "Error committing transaction", StatusCode: http.StatusInternalServerError}
	}

	return true, nil
}

func (ws *walletService) GetBalance(userID uint) (float64, *utils.HttpError) {
	wallet, err := ws.dao.NewWalletQuery().GetWalletByUserID(userID)
	if err != nil {
		log.Printf("Error retrieving wallet: %v", err)
		return 0, &utils.HttpError{Message: "Error retrieving wallet", StatusCode: http.StatusInternalServerError}
	}

	if wallet == nil {
		return 0, &utils.HttpError{Message: "Wallet not found", StatusCode: http.StatusNotFound}
	}

	return wallet.Balance, nil
}

func (ws *walletService) GetTransactions(userID uint) ([]datastruct.Transaction, *utils.HttpError) {
	wallet, err := ws.dao.NewWalletQuery().GetWalletByUserID(userID)
	if err != nil {
		log.Printf("Error retrieving wallet: %v", err)
		return nil, &utils.HttpError{Message: "Error retrieving wallet", StatusCode: http.StatusInternalServerError}
	}

	if wallet == nil {
		return nil, &utils.HttpError{Message: "Wallet not found", StatusCode: http.StatusNotFound}
	}

	transactions, err := ws.transactionService.GetTransactionsByWalletID(wallet.UserID)
	if err != nil {
		log.Printf("Error retrieving transactions: %v", err)
		return nil, &utils.HttpError{Message: "Error retrieving transactions", StatusCode: http.StatusInternalServerError}
	}

	return transactions, nil
}
