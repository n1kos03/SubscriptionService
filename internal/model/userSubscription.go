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
