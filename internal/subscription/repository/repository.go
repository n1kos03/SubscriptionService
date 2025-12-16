package repository

import (
	"SubscriptionService/internal/model"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, sub *model.UserSubscription) error
	GetSubByID(ctx context.Context, id uuid.UUID) (*model.UserSubscription, error)
	GetAllSubscriptions(ctx context.Context) ([]model.UserSubscription, error)
	ListUserSubs(ctx context.Context, user_id uuid.UUID) ([]model.UserSubscription, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, sub *model.UserSubscription) error
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

func (r *pgRepository) GetSubByID(ctx context.Context, id uuid.UUID) (*model.UserSubscription, error) {
	var sub model.UserSubscription
	err := r.db.WithContext(ctx).Model(&model.UserSubscription{}).First(&sub, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &sub, nil
}

func (r *pgRepository) GetAllSubscriptions(ctx context.Context) ([]model.UserSubscription, error) {
	var subs []model.UserSubscription
	err := r.db.WithContext(ctx).Model(&model.UserSubscription{}).Find(&subs).Error
	if err != nil {
		return nil, err
	}

	return subs, nil
}

func (r *pgRepository) ListUserSubs(ctx context.Context, user_id uuid.UUID) ([]model.UserSubscription, error) {
	var subs []model.UserSubscription
	err := r.db.WithContext(ctx).Model(&model.UserSubscription{}).Where("user_id = ?", user_id).Find(&subs).Error
	if err != nil {
		return nil, err
	}

	return subs, nil
}

func (r *pgRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&model.UserSubscription{}, "id = ?", id).Error
}

func (r *pgRepository) Update(ctx context.Context, sub *model.UserSubscription) error {
	err := r.db.WithContext(ctx).Model(&model.UserSubscription{}).Where("id = ?", sub.ID).Updates(sub).Error
	if err != nil {
		return err
	}

	return nil
}
