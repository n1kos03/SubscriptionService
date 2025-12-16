package service

import (
	"SubscriptionService/internal/model"
	"SubscriptionService/internal/subscription/repository"
	"SubscriptionService/pkg"
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"
)

type Service interface {
	CreateSubscription(ctx context.Context, sub *model.ServiceUserSubscription) (*model.UserSubscription, error)
	GetSubscriptionByID(ctx context.Context, id string) (*model.UserSubscription, error)
	GetAllSubscriptions(ctx context.Context) ([]model.UserSubscription, error)
	GetUserSubscriptions(ctx context.Context, user_id string) ([]model.UserSubscription, error)
	DeleteSubscription(ctx context.Context, id string) error
	UpdateSubscription(ctx context.Context, sub *model.ServiceUpdateUserSubscription) error
}

type service struct {
	rep repository.Repository
	log *slog.Logger
}

func InitService(rep repository.Repository, log *slog.Logger) Service {
	return &service{rep: rep, log: log}
}

func (s *service) CreateSubscription(ctx context.Context, sub *model.ServiceUserSubscription) (*model.UserSubscription, error) {
	if sub.Price < 0 {
		return nil, errors.New("price must be more or equal to 0")
	}

	if sub.ServiceName == "" {
		return nil, errors.New("service name can not be empty")
	}

	uid, err := uuid.Parse(sub.UserID)
	if err != nil {
		return nil, err
	}

	date, err := pkg.ParseDate(sub.StartDate)
	if err != nil {
		return nil, err
	}

	s.log.Info("Creating subscription", "user_id", sub.UserID, "service", sub.ServiceName)

	newSub := &model.UserSubscription{
		ID: uuid.New(),
		UserId: uid,
		ServiceName: sub.ServiceName,
		Price: sub.Price,
		StartDate: date,
	}

	err = s.rep.Create(ctx, newSub)
	if err != nil {
		s.log.Error("error on creating new subscription in database", "err", err)
		return nil, err
	}

	return newSub, nil
}

func (s *service) GetSubscriptionByID(ctx context.Context, id string) (*model.UserSubscription, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		s.log.Error("error while parsing id", "err", err)
		return nil, err
	}

	return s.rep.GetSubByID(ctx, uid)
}

func (s *service) GetAllSubscriptions(ctx context.Context) ([]model.UserSubscription, error) {
	return s.rep.GetAllSubscriptions(ctx)
}

func (s *service) GetUserSubscriptions(ctx context.Context, user_id string) ([]model.UserSubscription, error) {
	uid, err := uuid.Parse(user_id) 
	if err != nil {
		s.log.Error("error while parsing user_id", "err", err)
		return nil, err
	}

	return s.rep.ListUserSubs(ctx, uid)
}

func (s *service) DeleteSubscription(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		s.log.Error("error while parsing id", "err", err)
		return err
	}

	return s.rep.Delete(ctx, uid)
}

func (s *service) UpdateSubscription(ctx context.Context, sub *model.ServiceUpdateUserSubscription) error {
	var uSub model.UserSubscription
	
	id, err := uuid.Parse(sub.ID)
	if err != nil {
		s.log.Error("error while parsing id", "err", err)
		return err
	}
	uSub.ID = id

	uid, err := uuid.Parse(sub.UserId)
	if err != nil {
		s.log.Error("error while parsing id", "err", err)
		return err
	}
	uSub.UserId = uid

	if sub.Price != nil {
		uSub.Price = *sub.Price
	}

	if sub.ServiceName != nil {
		uSub.ServiceName = *sub.ServiceName
	}

	if sub.StartDate != nil {
		uSub.StartDate, err = pkg.ParseDate(*sub.StartDate)
		if err != nil {
			s.log.Error("error while parsing start date of subscription", "err", err)
			return err
		}
	}

	// uSub.StartDate, _ = time.Parse(time.RFC3339, *sub.StartDate)

	return s.rep.Update(ctx, &uSub)
}
