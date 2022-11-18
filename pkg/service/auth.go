package service

import (
	"crypto/sha1"
	"fmt"
	"time"
	balance "github.com/RomanMRR/user-balance"
	"github.com/RomanMRR/user-balance/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt = "wefwkevwelkv32r23vgref32fewv"
	signingKey = "wefervernvbobinroeibno4b"
	tokenTTL = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user balance.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}