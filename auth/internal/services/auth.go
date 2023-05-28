package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

const (
	secretKey  = "secret"
	expiration = 24 * time.Hour
)

type tokenCLaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
}

func NewAuthService() *authService {
	return &authService{}
}

type authService struct{}

func (s *authService) GetUserJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &tokenCLaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(expiration)),
			IssuedAt:  jwt.At(time.Now()),
		},
		Email: email,
	})
	return token.SignedString([]byte(secretKey))
}

func (s *authService) GetUserEmailByJWT(raw string) (string, error) {
	token, err := jwt.ParseWithClaims(raw, &tokenCLaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", fmt.Errorf("cannot parse jwt: %v", err)
	}

	parsed := token.Claims.(*tokenCLaims)

	if parsed.ExpiresAt.Time.Unix() < time.Now().Unix() {
		return "", errors.New("token expired")
	}

	return token.Claims.(*tokenCLaims).Email, nil
}
