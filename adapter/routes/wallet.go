package routes

import (
	"github.com/NicholasLiem/Paper_BE_Test/adapter/structs"
	"github.com/NicholasLiem/Paper_BE_Test/internal/app"
	"github.com/gin-gonic/gin"
)

func WalletRoutes(server app.MicroserviceServer) structs.RoutePrefix {
	return structs.RoutePrefix{
		Prefix: "/api/wallets",
		SubRoutes: []structs.Route{
			{
				"Top up wallet",
				"POST",
				"/topup",
				server.Topup,
				true,
				[]gin.HandlerFunc{},
			},
			{
				"Withdraw from wallet",
				"POST",
				"/withdraw",
				server.Withdraw,
				true,
				[]gin.HandlerFunc{},
			},
			{
				"Get wallet balance by user ID",
				"GET",
				"/balance/:user_id",
				server.GetBalance,
				false,
				[]gin.HandlerFunc{},
			},
		},
	}
}
