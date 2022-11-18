package handler

import (
	"github.com/RomanMRR/user-balance/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("/api")
	{
		balance := api.Group("/balance")
		{
			// balance.POST("/")
			balance.GET("/", h.getWalletById)
			balance.PUT("/", h.updateWallet)
			
			// balance.DELETE("/:id", )
		}
		transaction := balance.Group("/reserve")
		{
			transaction.PUT("/", h.addReserveWallet)
		}
		transactionSuccessful := balance.Group("/withdraw")
		{
			transactionSuccessful.PUT("/", h.withrawReserveWallet)
		}
	}

	return router
}
