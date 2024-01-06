package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/YnMann/chat_backend/internal/auth"
	"github.com/YnMann/chat_backend/internal/models"
	"github.com/golang-jwt/jwt"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *models.User `json:"user"`
}

type AuthUseCase struct {
	userRepo       auth.AuthRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUseCase(
	userRepo auth.AuthRepository,
	hashSalt string,
	signingKey []byte,
	tokenTTLSeconds time.Duration,
) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTLSeconds,
	}
}

func (a *AuthUseCase) SignUp(ctx context.Context, data *models.User) error {
	pwd := sha1.New()
	pwd.Write([]byte(data.Password))
	pwd.Write([]byte(a.hashSalt))

	user := &models.User{
		Email:     data.Email,
		Username:  data.Username,
		Password:  fmt.Sprintf("%x", pwd.Sum(nil)),
		FirstName: data.FirstName,
		LastName:  data.LastName,
	}

	return a.userRepo.CreateUser(ctx, user)
}

func (a *AuthUseCase) SignIn(ctx context.Context, username string, password string) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := a.userRepo.GetUser(ctx, username, password)
	if err != nil {
		return "", auth.ErrInvalidAccessToken
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(a.expireDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signingKey)
}

func (a *AuthUseCase) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&AuthClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return a.signingKey, nil
		})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, auth.ErrInvalidAccessToken
}
