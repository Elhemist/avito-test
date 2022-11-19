package handler

import (
	"fmt"
	"net/http"

	atest "github.com/Elhemist/avito-test/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateOrder(c *gin.Context) {
	var input atest.Order
	if err := c.BindJSON(&input); err != nil {

		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Order.CreateOrder(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) RevenueOrder(c *gin.Context) {
	var input atest.Order
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("nya1")
	id, err := h.services.Order.RevenueOrder(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
