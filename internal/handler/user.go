package handler

import (
	"net/http"

	atest "github.com/Elhemist/avito-test/models"
	"github.com/gin-gonic/gin"
)

const (
	userCtx = "userId"
)

// @Summary CheckBalance
// @Tags user
// @Description Check user's balance
// @ID check-order
// @Accept  json
// @Produce  json
// @Param input body atest.User true "User info"
// @Success 200 {object} atest.User
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/ [get]
func (h *Handler) CheckBalance(c *gin.Context) {
	var input atest.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.services.User.CheckUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"userId":  user.Id,
		"Balance": user.Balance,
	})
}

// @Summary UpdateBalance
// @Tags user
// @Description Add some cash to user's balance
// @ID check-order
// @Accept  json
// @Produce  json
// @Param input body atest.User true "User info"
// @Success 200 {object} atest.User
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/ [post]
func (h *Handler) UpdateBalance(c *gin.Context) {
	var input atest.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.services.User.UpdateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"userId":  user.Id,
		"Balance": user.Balance,
	})

}
