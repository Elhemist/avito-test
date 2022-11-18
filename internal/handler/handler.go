package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	user := router.Group("/user")
	{
		user.POST("/accrual", h.UpdateBalance)
		user.GET("/:id", h.CheckBalance)
	}
	product := router.Group("/order")
	{
		product.POST("/reserve", h.CreateOrder)
		product.POST("/revenue", h.RevenueOrder)
	}

	return router
}
