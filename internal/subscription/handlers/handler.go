package handlers

import (
	"SubscriptionService/internal/model"
	"SubscriptionService/internal/subscription/service"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	srv service.Service
	log *slog.Logger
}

func NewRouter(h *Handler) *gin.Engine {
	router := gin.Default()

	sub := router.Group("/")
	{
		sub.POST("subscription/", h.CreateSubscription)
	}

	return router
}

func InitHandler(s service.Service, log *slog.Logger) *Handler {
	return &Handler{srv: s, log: log}
}

func (h *Handler) CreateSubscription(c *gin.Context) {
	if c.Request.Method == http.MethodPost {
		var sub model.ServiceUserSubscription
		if err := c.ShouldBindBodyWithJSON(&sub); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		createdSub, err := h.srv.CreateSubscription(c.Request.Context(), &sub); 
		if err != nil {
			h.log.Error("error while creating subscription", "err", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id": createdSub.ID.String(),
		})
	}
}