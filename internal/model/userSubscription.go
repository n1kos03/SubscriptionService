package model

import (
	"time"

	"github.com/google/uuid"
)

type UserSubscription struct {
	ID          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"user_id"`
	ServiceName string    `json:"service_name"`
	Price       int       `json:"price"`
	StartDate   time.Time `json:"start_date"`
}


type ServiceUserSubscription struct {
	UserID string	`json:"user_id"`
	ServiceName string `json:"service_name"`
	Price int `json:"price"`
	StartDate string `json:"start_date"`
}

type ServiceUpdateUserSubscription struct {
	ID          string `json:"id"`
	UserId      string `json:"user_id"`
	ServiceName *string    `json:"service_name"`
	Price       *int       `json:"price"`
	StartDate   *string `json:"start_date"`
}

type SummarySubData struct {
	UserID *string `json:"user_id"`
	ServiceName *string `json:"service_name"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
}
