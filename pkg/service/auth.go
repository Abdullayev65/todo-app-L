package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/AbdullohAbdullayev/todo-app-L.git"
	"github.com/AbdullohAbdullayev/todo-app-L.git/pkg/repository"
	"github.com/AbdullohAbdullayev/todo-app-L.git/pkg/utill"
	"time"
)

const (
	salt           = "kmedytnzbdewkl"
	signingKey     = "vsdgvtrbvsdsac btropfclbr"
	tokenValidTime = 72 * time.Hour
)

type AuthService struct {
	repo     *repository.Repository
	tokenJWT *utill.TokenJWT
}

func newAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{
		repo:     repo,
		tokenJWT: utill.NewToken(signingKey, tokenValidTime),
	}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.hashingStr(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) hashingStr(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateTokenIfExists(username, password string) (string, error) {
	userId, err := s.repo.UserIdBy(username, s.hashingStr(password))
	if err != nil {
		return "", err
	}
	return s.tokenJWT.GenerateToken(userId)
}

func (s *AuthService) ParseToken(tokenStr string) (int, error) {
	return s.tokenJWT.ParseToken(tokenStr)
}
