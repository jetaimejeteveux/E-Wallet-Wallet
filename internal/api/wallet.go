package api

import (
	"ewallet-framework-1/constants"
	"ewallet-framework-1/helpers"
	"ewallet-framework-1/internal/interfaces"
	"ewallet-framework-1/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	WalletService interfaces.IWalletService
}

func (h *WalletHandler) Create(c *gin.Context) {
	var (
		req models.Wallet
		log = helpers.Logger
		ctx = c.Request.Context()
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("error parsing body : ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedParseRequest, nil)
	}

	if req.UserID == 0 {
		log.Error("user id is empty")
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedParseRequest, nil)
	}

	err := h.WalletService.Create(ctx, &req)
	if err != nil {
		log.Error("failed to create wallet : ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrFailedCreateWallet, nil)
	}

	helpers.SendResponseHTTP(c, http.StatusCreated, constants.SuccessMessage, req)
}
