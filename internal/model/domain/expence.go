package domain

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	Expenditure_type string
	Date             time.Time
	Amount           int64 // в копейках
	UserID           int64
}
