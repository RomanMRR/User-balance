package service

import (
	balance "github.com/RomanMRR/user-balance"
	"github.com/RomanMRR/user-balance/pkg/repository"
)

type Authorization interface {
	CreateUser(user balance.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type BalanceWallet interface {
	// Update(userId int, wallet balance.Wallet) (int, error)
	GetWallet(userId int)(balance.Wallet, error)
}

type Service struct {
	Authorization
	BalanceWallet
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		BalanceWallet: NewBalanceWalletService(repos.BalanceWallet),
	}
}