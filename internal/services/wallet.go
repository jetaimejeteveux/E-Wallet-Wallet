package services

import (
	"context"
	"ewallet-framework-1/internal/interfaces"
	"ewallet-framework-1/internal/models"
)

type WalletService struct {
	WalletRepo interfaces.IWalletRepository
}

func (s *WalletService) Create(ctx context.Context, wallet *models.Wallet) error {
	return s.WalletRepo.CreateWallet(ctx, wallet)
}
