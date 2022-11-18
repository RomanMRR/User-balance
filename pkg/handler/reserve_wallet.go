package handler

import (
	"fmt"
	"net/http"

	balance "github.com/RomanMRR/user-balance"
	"github.com/gin-gonic/gin"
)

type ReserveWalletInput struct {
	UserId       int  `json:"user_id" binding:"required"`
	OrderId       int  `json:"order_id" binding:"required"`
	ServiceId	int	`json:"service_id" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}


// @Summary Reserve funds
// @Tags reserve wallet
// @Description Reserving money from the main account
// @ID reserve-funds
// @Accept  json
// @Produce  json
// @Param input body balance.UpdateReserveWallet true "How many funds to reserve"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/balance/withdraw [put]
func (h *Handler) addReserveWallet(c *gin.Context) {
	var input balance.UpdateReserveWallet
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	wallet, err := h.services.BalanceWallet.GetWallet(*input.User_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if wallet.Amount < *input.Amount {
		println(*input.Amount)
		println(wallet.Amount)
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Insufficient funds"))
		return
	}

	if err := h.services.AddAmount(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}


	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Debiting funds
// @Tags reserve wallet
// @Description Debiting funds from a reserve wallet
// @ID debiting-funds
// @Accept  json
// @Produce  json
// @Param input body balance.UpdateReserveWallet true "How many funds to withdraw"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/balance/reserve [put]
func (h *Handler) withdrawReserveWallet(c *gin.Context) {
	var input balance.UpdateReserveWallet
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}


	if err := h.services.WithdrawAmount(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}


	c.JSON(http.StatusOK, statusResponse{"ok"})
}