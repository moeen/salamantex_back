package server

import (
	"gihub.com/moeen/salamantex_back/controllers/account"
	"gihub.com/moeen/salamantex_back/controllers/currency"
	"gihub.com/moeen/salamantex_back/controllers/transaction"
	"gihub.com/moeen/salamantex_back/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	api := router.Group("/api/v1")

	{
		api.POST("/user/register", account.Register)
		api.POST("/user/login", account.Login)
	}

	{
		auth := api.Group("/")
		auth.Use(middlewares.AuthenticationNeeded())

		auth.GET("/user", account.GetUser)

		auth.POST("/currency/add", currency.AddCurrency)

		auth.POST("/tx/send", transaction.Send)
		auth.GET("/tx/history", transaction.History)
		auth.GET("/tx/state/:id", transaction.Info)
	}

	return router
}
