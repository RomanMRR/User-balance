package handler

import (
	"github.com/RomanMRR/user-balance/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/files" // swagger embed files

	_ "github.com/RomanMRR/user-balance/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("/api")
	{
		balance := api.Group("/balance")
		{
			balance.GET("/", h.getWalletById)
			balance.PUT("/", h.updateWallet)
			
		}
		transaction := balance.Group("/reserve")
		{
			transaction.PUT("/", h.addReserveWallet)
		}
		transactionSuccessful := balance.Group("/withdraw")
		{
			transactionSuccessful.PUT("/", h.withdrawReserveWallet)
		}
	}

	return router
}
