package service

import (
	balance "github.com/RomanMRR/user-balance"
	"github.com/RomanMRR/user-balance/pkg/repository"
)

type Authorization interface {
	CreateUser(user balance.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Balance interface {

}

type Service struct {
	Authorization
	Balance
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}