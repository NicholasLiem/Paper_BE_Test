package repository

import (
	"errors"
	"github.com/NicholasLiem/Paper_BE_Test/internal/datastruct"
	"gorm.io/gorm"
)

type WalletQuery interface {
	BeginTransaction() *gorm.DB
	CreateWallet(wallet datastruct.Wallet) (bool, error)
	CreateWalletTx(wallet datastruct.Wallet, tx *gorm.DB) (bool, error)
	UpdateWalletTx(walletID uint, updatedWallet datastruct.Wallet, tx *gorm.DB) (bool, error)
	DeleteWallet(walletID uint) (bool, error)
	GetWallet(walletID uint) (*datastruct.Wallet, error)
	GetWalletByUserID(userID uint) (*datastruct.Wallet, error)
}

type walletQuery struct {
	pgdb *gorm.DB
}

func NewWalletQuery(db *gorm.DB) WalletQuery {
	return &walletQuery{pgdb: db}
}

func (wq *walletQuery) BeginTransaction() *gorm.DB {
	return wq.pgdb.Begin()
}

func (wq *walletQuery) CreateWallet(wallet datastruct.Wallet) (bool, error) {
	result := wq.pgdb.Create(&wallet)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (wq *walletQuery) CreateWalletTx(wallet datastruct.Wallet, tx *gorm.DB) (bool, error) {
	result := tx.Create(&wallet)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (wq *walletQuery) DeleteWallet(walletID uint) (bool, error) {
	result := wq.pgdb.Delete(&datastruct.Wallet{}, walletID)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (wq *walletQuery) UpdateWalletTx(walletID uint, updatedWallet datastruct.Wallet, tx *gorm.DB) (bool, error) {
	var wallet datastruct.Wallet
	if err := tx.First(&wallet, walletID).Error; err != nil {
		return false, err
	}

	result := tx.Model(&wallet).Updates(updatedWallet)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (wq *walletQuery) GetWallet(walletID uint) (*datastruct.Wallet, error) {
	var wallet datastruct.Wallet
	result := wq.pgdb.First(&wallet, walletID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &wallet, nil
}

func (wq *walletQuery) GetWalletByUserID(userID uint) (*datastruct.Wallet, error) {
	var wallet datastruct.Wallet
	if err := wq.pgdb.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &wallet, nil
}
