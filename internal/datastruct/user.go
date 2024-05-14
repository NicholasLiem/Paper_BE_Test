package datastruct

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:100;not null;unique" json:"name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Wallet    Wallet    `gorm:"foreignKey:UserID"`
}

func (user *User) BeforeSave(db *gorm.DB) (err error) {
	if len(user.Name) < 3 {
		return errors.New("the name must be at least 3 characters long")
	}
	if len(user.Email) < 5 {
		return errors.New("the email must be at least 5 characters long")
	}
	return nil
}
