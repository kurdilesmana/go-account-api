// internal/domain/account.go
package domain

import "time"

type Transaction struct {
	ID              uint      `gorm:"primaryKey"`
	TransactionTime time.Time `json:"transaction_time"`
	TransactionCode string    `json:"transaction_code"`
	Amount          float64   `json:"amount"`
}

type TransactionDetail struct {
	ID            uint    `gorm:"primaryKey"`
	TransactionID int     `json:"transaction_id"`
	Mutation      string  `json:"mutation"`
	AccountNumber string  `json:"account_number"`
	Amount        float64 `json:"amount"`
}

type TransactionMutations struct {
	TransactionTime time.Time `json:"transaction_time"`
	TransactionCode string    `json:"transaction_code"`
	Mutation        string    `json:"mutation"`
	Amount          float64   `json:"amount"`
}

type Saving struct {
	AccountNumber string  `json:"account_number" validate:"required"`
	Amount        float64 `json:"amount" validate:"required"`
}

type CashWithdrawl struct {
	AccountNumber string  `json:"account_number" validate:"required"`
	Amount        float64 `json:"amount" validate:"required"`
}

type Transfer struct {
	SrcAccountNumber  string  `json:"src_account_number" validate:"required"`
	DestAccountNumber string  `json:"dest_account_number" validate:"required"`
	Amount            float64 `json:"amount" validate:"required"`
}
