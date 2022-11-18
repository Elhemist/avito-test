package handler

import (
	"net/http"

	atest "github.com/Elhemist/avito-test/models"
	"github.com/gin-gonic/gin"
)

const (
	userCtx = "userId"
)

func (h *Handler) UpdateBalance(c *gin.Context) {

}

func (h *Handler) CheckBalance(c *gin.Context) {
	var input atest.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.User.CheckUser(input)
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
