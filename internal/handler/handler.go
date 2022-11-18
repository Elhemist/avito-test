package handler

import (
	"github.com/Elhemist/avito-test/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	user := router.Group("/user")
	{
		user.POST("/:id", h.UpdateBalance)
		user.GET("/:id", h.CheckBalance)
	}
	product := router.Group("/order")
	{
		product.GET("/reserve", h.CreateOrder)
		product.POST("/revenue", h.RevenueOrder)
	}

	return router
}
