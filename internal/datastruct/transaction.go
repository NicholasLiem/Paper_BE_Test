package datastruct

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type TransactionType string

const (
	CashIn  TransactionType = "cash_in"
	CashOut TransactionType = "cash_out"
)

type Transaction struct {
	ID        uint            `gorm:"primaryKey;autoIncrement"`
	WalletID  uint            `gorm:"not null"`
	Amount    float64         `gorm:"not null"`
	Type      TransactionType `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time       `gorm:"autoCreateTime"`
}

func (transaction *Transaction) BeforeSave(tx *gorm.DB) (err error) {
	if transaction.Type != CashIn && transaction.Type != CashOut {
		return errors.New("invalid transaction type")
	}
	return nil
}
