package adapter

import (
	"github.com/NicholasLiem/Paper_BE_Test/adapter/middleware"
	"github.com/NicholasLiem/Paper_BE_Test/adapter/routes"
	"github.com/NicholasLiem/Paper_BE_Test/adapter/structs"
	"github.com/NicholasLiem/Paper_BE_Test/internal/app"
	"github.com/gin-gonic/gin"
)

func NewRouter(server app.MicroserviceServer) *gin.Engine {
	router := gin.Default()

	structs.AppRoutes = append(structs.AppRoutes, routes.TransactionRoutes(server), routes.UserRoutes(server), routes.WalletRoutes(server))
	for _, routePrefix := range structs.AppRoutes {
		group := router.Group(routePrefix.Prefix)

		for _, route := range routePrefix.SubRoutes {
			routeGroup := group.Group(route.Pattern)

			if route.JSONRequest {
				routeGroup.Use(middleware.ApplyJSONMiddleware)
			}

			routeGroup.Use(route.Middleware...)

			switch route.Method {
			case "GET":
				routeGroup.GET("", route.HandlerFunc)
			case "POST":
				routeGroup.POST("", route.HandlerFunc)
			case "PATCH":
				routeGroup.PATCH("", route.HandlerFunc)
			case "DELETE":
				routeGroup.DELETE("", route.HandlerFunc)
			case "PUT":
				routeGroup.PUT("", route.HandlerFunc)
			}
		}
	}
	return router
}
