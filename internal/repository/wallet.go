package repository

import (
	"context"
	"ewallet-framework-1/internal/models"

	"gorm.io/gorm"
)

type WalletRepository struct {
	DB *gorm.DB
}

func (r *WalletRepository) CreateWallet(ctx context.Context, wallet *models.Wallet) error {
	return r.DB.Create(&wallet).Error
}
