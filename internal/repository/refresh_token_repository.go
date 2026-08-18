package repository

import (
	"github.com/aditya-singh-finbox/todo-api/internal/model"
	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		db: db,
	}
}

func (r *RefreshTokenRepository) Create(
	refreshToken *model.RefreshToken,
) error {
	return r.db.Create(refreshToken).Error
}

func (r *RefreshTokenRepository) GetByToken(
	token string,
) (*model.RefreshToken, error) {
	var refreshToken model.RefreshToken

	err := r.db.Where("token = ?", token).First(&refreshToken).Error

	if err != nil {
		return nil, err
	}

	return &refreshToken, nil
}

func (r *RefreshTokenRepository) Delete(token string) error {
	return r.db.Where("token = ?", token).Delete(&model.RefreshToken{}).Error
}
