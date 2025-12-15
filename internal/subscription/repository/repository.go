package repository

import (
	"SubscriptionService/internal/model"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, sub *model.UserSubscription) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.UserSubscription, error)
	GetAllSubscriptions(ctx context.Context) ([]model.UserSubscription, error)
	ListUserSubs(ctx context.Context, user_id uuid.UUID) ([]model.UserSubscription, error)
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

func (r *pgRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.UserSubscription, error) {
	sub, err := gorm.G[model.UserSubscription](r.db).Where("id = ?", id).First(ctx)
	
	return &sub, err
}

func (r *pgRepository) GetAllSubscriptions(ctx context.Context) ([]model.UserSubscription, error) {
	return gorm.G[model.UserSubscription](r.db).Find(ctx)
}

func (r *pgRepository) ListUserSubs(ctx context.Context, user_id uuid.UUID) ([]model.UserSubscription, error) {
	return gorm.G[model.UserSubscription](r.db).Where("user_id = ?", user_id).Find(ctx)
}
