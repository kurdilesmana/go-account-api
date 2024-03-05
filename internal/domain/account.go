// internal/domain/account.go
package domain

type Account struct {
	AccountNumber string  `gorm:"primaryKey" json:"account_number" validate:"required"`
	Name          string  `json:"name" validate:"required"`
	NIK           string  `json:"nik" validate:"required"`
	PhoneNumber   string  `json:"phone_number" validate:"required"`
	PIN           string  `json:"pin" validate:"required"`
	Balance       float64 `json:"balance"`
}

type CreateAccount struct {
	Name        string `json:"name" validate:"required"`
	NIK         string `json:"nik" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	PIN         string `json:"pin" validate:"required"`
}
