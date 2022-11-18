package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	user := router.Group("/user")
	{
		user.POST("/accrual")
		user.POST("/reserve")
		user.GET("/:id")
	}
	product := router.Group("/order")
	{
		product.POST("/revenue")
	}

	return router
}
