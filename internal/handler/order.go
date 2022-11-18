package handler

import (
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

	// x := input.Id
	// fmt.Printf(fmt.Sprint(x))
}

func (h *Handler) RevenueOrder(c *gin.Context) {

}
