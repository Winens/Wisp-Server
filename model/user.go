package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID     int64  `gorm:"primaryKey"`
	AuthID string `gorm:"unique;not null"`

	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func (u *User) GeneratePasswordDigest(password string) error {
	digest, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(digest)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
