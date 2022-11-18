package service

import (
	balance "github.com/RomanMRR/user-balance"
	"github.com/RomanMRR/user-balance/pkg/repository"
)

type Authorization interface {
	CreateUser(user balance.User) (int, error)
}

type BalanceWallet interface {
	Update(input balance.UpadateWallet) error
	GetWallet(userId int)(balance.Wallet, error)
}

type ReserveWallet interface {
	AddAmount(input balance.UpdateReserveWallet) error
	WithdrawAmount(input balance.UpdateReserveWallet) error
}

type Service struct {
	Authorization
	BalanceWallet
	ReserveWallet
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		BalanceWallet: NewBalanceWalletService(repos.BalanceWallet),
		ReserveWallet: NewReserveWalletService(repos.ReserveWallet),
	}
}