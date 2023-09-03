package model

import "time"

type PublicKey struct {
	ID     int64  `json:"id" gorm:"primaryKey"`
	UserID int64  `json:"userID"`
	Key    []byte `json:"key"`

	CreatedAt time.Time `json:"createdAt"`
}
