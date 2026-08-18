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
) (string, string, error) {

	token, err := s.refreshTokenRepo.GetByToken(refreshToken)

	if err != nil {
		return "", "", errors.New("invalid refresh token")
	}

	if time.Now().After(token.ExpiresAt) {

		_ = s.refreshTokenRepo.Delete(refreshToken)

		return "", "", errors.New("refresh token expired")
	}

	// Generate new access token
	accessToken, err := s.jwtService.GenerateToken(token.UserID)

	if err != nil {
		return "", "", err
	}

	// Generate new refresh token
	newRefreshToken, err := auth.GenerateRefreshToken()

	if err != nil {
		return "", "", err
	}

	// Delete old refresh token
	err = s.refreshTokenRepo.Delete(refreshToken)

	if err != nil {
		return "", "", err
	}

	// Store new refresh token
	newRefreshTokenModel := &model.RefreshToken{
		Token:     newRefreshToken,
		UserID:    token.UserID,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	err = s.refreshTokenRepo.Create(newRefreshTokenModel)

	if err != nil {
		return "", "", err
	}

	return accessToken, newRefreshToken, nil
}

func (s *AuthService) Logout(refreshToken string) error {
	return s.refreshTokenRepo.Delete(refreshToken)
}
