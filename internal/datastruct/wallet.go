package datastruct

type Wallet struct {
	UserID  uint    `gorm:"primaryKey"`
	Balance float64 `gorm:"default:0"`
	User    *User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
