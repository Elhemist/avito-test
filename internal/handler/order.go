package handler

import (
	"fmt"
	"net/http"

	atest "github.com/Elhemist/avito-test/models"
	"github.com/gin-gonic/gin"
)

// @Summary CreateOrder
// @Tags order
// @Description Create order
// @ID create-order
// @Accept  json
// @Produce  json
// @Param input body atest.OrderResp true "order info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /order/reserve [post]
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

// @Summary RevenueOrder
// @Tags order
// @Description Revenue order
// @ID revenue-order
// @Accept  json
// @Produce  json
// @Param input body atest.OrderResp true "revenue info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /order/revenue [post]
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
