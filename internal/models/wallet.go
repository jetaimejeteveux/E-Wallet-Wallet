package models

import "time"

type Wallet struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id" gorm:"column:user_id"`
	Balance   float64   `json:"balance" gorm:"balance;type:decimal(15,2)"`
	CreatedAt time.Time `json:"-" gorm:"column:created_at;type:timestamp"`
	UpdatedAt time.Time `json:"-" gorm:"column:updated_at;type:timestamp"`
}

func (*Wallet) TableName() string {
	return "wallets"
}

type WalletTransaction struct {
	ID                    int
	WalletID              int       `json:"wallet_id" gorm:"column:wallet_id"`
	Amount                float64   `json:"amount" gorm:"column:amount"`
	WalletTransactionType string    `json:"wallet_transaction_type" gorm:"column:walllet_transaction_type;type:ENUM('CREDIT', 'DEBIT')"`
	Reference             string    `gorm:"column:reference;type:varchar(100);unique"`
	CreatedAt             time.Time `json:"date" gorm:"column:created_at;type:timestamp"`
	UpdatedAt             time.Time `json:"-" gorm:"column:updated_at;type:timestamp"`
}

func (*WalletTransaction) TableName() string {
	return "wallet_transactions"
}
