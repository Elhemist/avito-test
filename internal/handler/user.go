package handler

import (
	"net/http"

	atest "github.com/Elhemist/avito-test/models"
	"github.com/gin-gonic/gin"
)

const (
	userCtx = "userId"
)

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
