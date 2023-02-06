package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID          int64
	DefaultCurrency string
}
