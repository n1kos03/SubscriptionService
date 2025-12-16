package model

import (
	"time"

	"github.com/google/uuid"
)

type UserSubscription struct {
	ID          uuid.UUID `json:"id" example:"7639ca88-80df-4d91-ae20-78ac3431ee11"`
	UserId      uuid.UUID `json:"user_id" example:"7639ca88-80df-4d91-ae20-78ac3431ee11"`
	ServiceName string    `json:"service_name" example:"Yandex Plus"`
	Price       int       `json:"price" example:"400"`
	StartDate   time.Time `json:"start_date" example:"07-2025"`
}

type ServiceUserSubscription struct {
	UserID      string `json:"user_id" example:"7639ca88-80df-4d91-ae20-78ac3431ee11"`
	ServiceName string `json:"service_name" example:"Yandex Plus"`
	Price       int    `json:"price" example:"400"`
	StartDate   string `json:"start_date" example:"07-2025"`
}

type ServiceUpdateUserSubscription struct {
	ID          string  `json:"id" example:"7639ca88-80df-4d91-ae20-78ac3431ee11"`
	UserId      string  `json:"user_id" example:"7639ca88-80df-4d91-ae20-78ac3431ee11"`
	ServiceName *string `json:"service_name" example:"Yandex Plus"`
	Price       *int    `json:"price" example:"400"`
	StartDate   *string `json:"start_date" example:"07-2025"`
}

type SummarySubData struct {
	UserID      *string `json:"user_id"`
	ServiceName *string `json:"service_name"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
}
