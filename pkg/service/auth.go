package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/AbdullohAbdullayev/todo-app-L.git"
	"github.com/AbdullohAbdullayev/todo-app-L.git/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt           = "kmedytnzbdewkl"
	signingKey     = "vsdgvtrbvsdsac btropfclbr"
	tokenValidTime = 72 * time.Hour
)

type AuthService struct {
	repo *repository.Repository
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
}

func newAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.hashingStr(user.Password)
	return s.repo.CreateUser(user)
}

func (s AuthService) hashingStr(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
func (s AuthService) GenerateTokenIfExists(username, password string) (string, error) {
	userId, err := s.repo.UserIdBy(username, s.hashingStr(password))
	if err != nil {
		return "", err
	}
	return generateToken(int64(userId), tokenValidTime, signingKey)
}

func generateToken(userId int64, interval time.Duration, signingKey string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodES256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(interval).Unix(),
		},
		UserId: userId,
	})
	return token.SignedString([]byte(signingKey))
}
