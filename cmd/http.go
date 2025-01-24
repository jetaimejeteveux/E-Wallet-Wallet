package cmd

import (
	"ewallet-framework-1/helpers"
	"ewallet-framework-1/internal/api"
	"ewallet-framework-1/internal/interfaces"
	"ewallet-framework-1/internal/repository"
	"ewallet-framework-1/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	d := dependencyInject()
	r := gin.Default()

	r.GET("/health", d.HealthcheckAPI.HealthcheckHandlerHTTP)
	walletV1 := r.Group("/wallet/v1")
	walletV1.POST("/", d.WalletHandler.Create)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	HealthcheckAPI interfaces.IHealthcheckHandler
	WalletHandler  interfaces.IWalletHandler
}

func dependencyInject() Dependency {
	healthcheckSvc := &services.Healthcheck{}
	healthcheckAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	walletRepo := &repository.WalletRepository{
		DB: helpers.DB,
	}

	walletSvc := &services.WalletService{
		WalletRepo: walletRepo,
	}
	walletHandler := &api.WalletHandler{
		WalletService: walletSvc,
	}

	return Dependency{
		HealthcheckAPI: healthcheckAPI,
		WalletHandler:  walletHandler,
	}
}
