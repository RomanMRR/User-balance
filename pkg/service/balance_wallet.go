package service

import (
	balance "github.com/RomanMRR/user-balance"
	"github.com/RomanMRR/user-balance/pkg/repository"
)

type BalanceWalletService struct {
	repo repository.BalanceWallet
}

func NewBalanceWalletService(repo repository.BalanceWallet) *BalanceWalletService {
	return &BalanceWalletService{repo: repo}
}

// func (s *BalanceWalletService) Update(userId int, wallet balance.Wallet) (int, error) {
// 	return s.repo.Update(userId, wallet)
// }

func (s *BalanceWalletService) GetWallet(userId int)(balance.Wallet, error) {
	return s.repo.GetWallet(userId)
}

