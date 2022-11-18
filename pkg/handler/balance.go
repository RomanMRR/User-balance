package handler

import (
	"net/http"

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

type UserInput struct {
	UserId int `json:"user_id"`
}

func (h *Handler) getWalletById(c *gin.Context) {
	var input UserInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	

	wallet, err := h.services.BalanceWallet.GetWallet(input.UserId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, wallet)
}

func (h *Handler) updateWallet(c *gin.Context) {
	var input balance.UpadateWallet
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Update(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}


	c.JSON(http.StatusOK, statusResponse{"ok"})
}