package handler

import (
	"net/http"
	"strconv"

	balance "github.com/RomanMRR/user-balance"
	"github.com/gin-gonic/gin"
)



type UserInput struct {
	UserId int `json:"user_id"`
}

// @Summary Get wallet
// @Tags wallet
// @Description get the user's wallet
// @ID get-wallet
// @Produce  json
// @Param  id path int true "User ID"
// @Success 200 {object} balance.Wallet
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/balance/{id} [get]
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

// @Summary Update wallet
// @Tags wallet
// @Description Updating funds on the user's account
// @ID update-wallet
// @Accept  json
// @Produce  json
// @Param input body balance.UpdateWallet true "How many funds to transfer"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/balance [put]
func (h *Handler) updateWallet(c *gin.Context) {
	var input balance.UpdateWallet
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