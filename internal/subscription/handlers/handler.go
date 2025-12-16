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
		sub.GET("subscription/:id/", h.GetSubscriptionByID)
		sub.GET("subscriptions/", h.GetAllSubscriptions)
		sub.GET("subscriptions/:user_id/", h.GetUserSubscriptions)
		sub.DELETE("subscription/:id/", h.DeleteSubscription)
		sub.PATCH("subscription/", h.UpdateSubscription)
		sub.GET("subscription/sum/", h.SummaryPriceSub)
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

func (h *Handler) GetSubscriptionByID(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		id := c.Param("id")

		sub, err := h.srv.GetSubscriptionByID(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"subscription": sub,
		})
	}	
}

func (h *Handler) GetAllSubscriptions(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		subs, err := h.srv.GetAllSubscriptions(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"subscriptions": subs,
		})
	}
}

func (h *Handler) GetUserSubscriptions(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		user_id := c.Param("user_id")
		uSubs, err := h.srv.GetUserSubscriptions(c.Request.Context(), user_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user_subs": uSubs,
		})
	}
}

func (h *Handler) DeleteSubscription(c *gin.Context) {
	if c.Request.Method == http.MethodDelete {
		id := c.Param("id")
		err := h.srv.DeleteSubscription(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Subscription succefully deleted"})
	}
}

func (h *Handler) UpdateSubscription(c *gin.Context) {
	if c.Request.Method == http.MethodPatch {
		var uSub model.ServiceUpdateUserSubscription
		if err := c.ShouldBindBodyWithJSON(&uSub); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		
		if err := h.srv.UpdateSubscription(c.Request.Context(), &uSub); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Subscription succefully updated",
			"updated subscription": uSub,
		})
	}
}

func (h *Handler) SummaryPriceSub(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		subInfo := model.SummarySubData{
			UserID: nil,
			ServiceName: nil,
			StartDate: c.Query("start_date"),
			EndDate: c.Query("end_date"),
		}

		if uid := c.Query("user_id"); uid != "" {
			subInfo.UserID = &uid
		}

		if serviceName := c.Query("service_name"); serviceName != "" {
			subInfo.ServiceName = &serviceName
		}

		total, err := h.srv.SummuryPriceSub(c.Request.Context(), subInfo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"total sum": total})
	}
}
