package service

import (
	"errors"
	"time"

	"github.com/aditya-singh-finbox/todo-api/internal/auth"
	"github.com/aditya-singh-finbox/todo-api/internal/model"
	"github.com/aditya-singh-finbox/todo-api/internal/repository"
)

type AuthService struct {
	userService      *UserService
	jwtService       *auth.JWTService
	refreshTokenRepo *repository.RefreshTokenRepository
}

func NewAuthService(
	userService *UserService,
	jwtService *auth.JWTService,
	refreshTokenRepo *repository.RefreshTokenRepository,
) *AuthService {
	return &AuthService{
		userService:      userService,
		jwtService:       jwtService,
		refreshTokenRepo: refreshTokenRepo,
	}
}

func (s *AuthService) Login(
	email string,
	password string,
) (string, string, *model.User, error) {
	user, err := s.userService.Login(email, password)

	if err != nil {
		return "", "", nil, err
	}

	accessToken, err := s.jwtService.GenerateToken(user.ID)

	if err != nil {
		return "", "", nil, err
	}

	refreshToken, err := auth.GenerateRefreshToken()

	if err != nil {
		return "", "", nil, err
	}

	refreshTokenModel := &model.RefreshToken{
		Token:     refreshToken,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	err = s.refreshTokenRepo.Create(refreshTokenModel)

	if err != nil {
		return "", "", nil, err
	}

	return accessToken, refreshToken, user, nil
}

func (s *AuthService) Refresh(
	refreshToken string,
) (string, error) {
	token, err := s.refreshTokenRepo.GetByToken(refreshToken)

	if err != nil {
		return "", errors.New("invalid refresh token")
	}
	if time.Now().After(token.ExpiresAt) {
		_ = s.refreshTokenRepo.Delete(refreshToken)
		return "", errors.New("refresh token expired")
	}

	accessToken, err := s.jwtService.GenerateToken(token.UserID)

	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (s *AuthService) Logout(refreshToken string) error {
	return s.refreshTokenRepo.Delete(refreshToken)
}
