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

func (s *BalanceWalletService) Update(input balance.UpadateWallet) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(input)
}

func (s *BalanceWalletService) GetWallet(userId int)(balance.Wallet, error) {
	return s.repo.GetWallet(userId)
}

