package handler

import (
	"net/http"
	"strconv"

	balance "github.com/RomanMRR/user-balance"
	"github.com/gin-gonic/gin"
)


func (h *Handler) updateBalanceById(c *gin.Context) {
	var input balance.Wallet
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

type getWalletResponse struct {
	Data balance.Wallet `json:"data"`
}

func (h *Handler) getWalletById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	wallet, err := h.services.BalanceWallet.GetWallet(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, wallet)
}