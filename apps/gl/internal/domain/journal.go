// internal/domain/account.go
package domain

import "time"

type Journal struct {
	ID              uint      `gorm:"primaryKey"`
	TransactionDate time.Time `json:"transaction_date" validate:"required"`
	CreditAccount   string    `json:"credit_account" validate:"required"`
	DebetAccount    string    `json:"debet_account" validate:"required"`
	CreditAmount    float64   `json:"credit_amount" validate:"required"`
	DebetAmount     float64   `json:"debet_amount" validate:"required"`
}

type CreateJournal struct {
	TransactionDate time.Time `json:"transaction_date" validate:"required"`
	CreditAccount   string    `json:"credit_account" validate:"required"`
	DebetAccount    string    `json:"debet_account" validate:"required"`
	CreditAmount    float64   `json:"credit_amount" validate:"required"`
	DebetAmount     float64   `json:"debet_amount" validate:"required"`
}
