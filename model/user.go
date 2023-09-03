package model

type User struct {
	ID     int64  `gorm:"primaryKey"`
	AuthID string `gorm:"unique;not null"`

	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}
