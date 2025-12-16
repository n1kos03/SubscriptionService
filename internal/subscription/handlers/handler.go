package handlers

import (
	"SubscriptionService/internal/model"
	"SubscriptionService/internal/subscription/service"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "SubscriptionService/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

func InitHandler(s service.Service, log *slog.Logger) *Handler {
	return &Handler{srv: s, log: log}
}

// Create subscription godoc
// @Summary create subscription
// @Description create new subscription for user
// @Tags subscription
// @Accept json
// @Produce json
// @Param subscription body model.ServiceUserSubscription true "Subscription data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscription/ [post]
func (h *Handler) CreateSubscription(c *gin.Context) {
	if c.Request.Method == http.MethodPost {
		var sub model.ServiceUserSubscription
		if err := c.ShouldBindBodyWithJSON(&sub); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		createdSub, err := h.srv.CreateSubscription(c.Request.Context(), &sub)
		if err != nil {
			h.log.Error("error while creating subscription", "err", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"id": createdSub.ID.String(),
		})
	}
}

// GetSubscriptionByID godoc
// @Summery Get subscription by ID
// @Tags subscription
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 200 {object} model.UserSubscription
// @Failure 400 {object} map[string]interface{}
// @Router /subscription/{id}/ [get]
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

// GetAllSubscriptions godoc
// @Summary Get all subscriptions
// @Tags subscription
// @Produce json
// @Success 200 {array} model.UserSubscription
// @Failure 500 {object} map[string]interface{}
// @Router /subscriptions/ [get]
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

// GetUserSubscriptions godoc
// @Summary Get subscriptions by user
// @Tags subscription
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {array} model.UserSubscription
// @Faulure 400 {object} map[string]interface{}
// @Router /subscriptions/{user_id}/ [get]
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

// DeleteSubscription godoc
// @Summary Delete subscription
// @Tags subscription
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /subscription/{id}/ [delete]
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

// UpdateSubscription godoc
// @Summary Update subscription
// @Description Partial update of subscription
// @Tags subscription
// @Accept json
// @Produce json
// @Param subscription body model.ServiceUpdateUserSubscription true "Update data"
// @Success  200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /subscription/ [patch]
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

// SummaryPriceSub godoc
// @Summary Get total subscription cost
// @Description Calculate total price for subscriptions in specified period by user or service name
// @Tags subscription
// @Produce json
// @Param user_id query string false "User ID"
// @Param service_name query string false "Service name"
// @Param start_date query string true "Start date (MM-YYYY)"
// @Param end_date query string true "End date (MM-YYYY)"
// @Success 200 {object} map[string]int
// @Failure 500 {object} map[string]interface{}
// @Router /subscription/sum/ [get]
func (h *Handler) SummaryPriceSub(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		subInfo := model.SummarySubData{
			UserID:      nil,
			ServiceName: nil,
			StartDate:   c.Query("start_date"),
			EndDate:     c.Query("end_date"),
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
