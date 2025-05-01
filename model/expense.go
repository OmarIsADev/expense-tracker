package model

import "time"

type Expense struct {
	ID          uint16
	Description string
	Amount      float64
	CreatedAt   time.Time
}
