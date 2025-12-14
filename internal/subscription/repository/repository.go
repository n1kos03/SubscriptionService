package repository

import (
	"SubscriptionService/internal/model"
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, sub *model.UserSubscription) error
}

type pgRepository struct {
	db *gorm.DB
}

func InitRepository(db *gorm.DB) Repository {
	return &pgRepository{db: db}
}

func (r *pgRepository) Create(ctx context.Context, sub *model.UserSubscription) error {
	return r.db.WithContext(ctx).Create(sub).Error
}
