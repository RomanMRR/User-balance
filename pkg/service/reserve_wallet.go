package service

import (
	"errors"

	balance "github.com/RomanMRR/user-balance"
	"github.com/RomanMRR/user-balance/pkg/repository"
)

type ReserveWalletService struct {
	repo repository.ReserveWallet
}

func NewReserveWalletService(repo repository.ReserveWallet) *ReserveWalletService {
	return &ReserveWalletService{repo: repo}
}

func (s *ReserveWalletService) AddAmount(input balance.UpdateReserveWallet) error {
	return s.repo.AddAmount(input)
}

func (s *ReserveWalletService) WithdrawAmount(input balance.UpdateReserveWallet) error {
	reserveWalet, err := s.repo.GetReserveWallet(*input.User_id)
	if err != nil {
		return err
	}

	if reserveWalet.Amount < *input.Amount {
		return errors.New("insufficient funds in the reserve account")
	}
	return s.repo.WithdrawAmount(input)
}

