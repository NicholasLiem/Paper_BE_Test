package routes

import (
	"github.com/NicholasLiem/Paper_BE_Test/adapter/structs"
	"github.com/NicholasLiem/Paper_BE_Test/internal/app"
	"github.com/gin-gonic/gin"
)

func TransactionRoutes(server app.MicroserviceServer) structs.RoutePrefix {
	return structs.RoutePrefix{
		Prefix: "/api/transactions",
		SubRoutes: []structs.Route{
			{
				"Get transaction by ID",
				"GET",
				"/:id",
				server.GetTransaction,
				false,
				[]gin.HandlerFunc{},
			},
			{
				"Get all transactions for a wallet by wallet ID",
				"GET",
				"/wallets/:wallet_id",
				server.GetTransactionsByWalletID,
				false,
				[]gin.HandlerFunc{},
			},
		},
	}
}
