package routes

import (
	"github.com/NicholasLiem/Paper_BE_Test/adapter/structs"
	"github.com/NicholasLiem/Paper_BE_Test/internal/app"
	"github.com/gin-gonic/gin"
)

func UserRoutes(server app.MicroserviceServer) structs.RoutePrefix {
	return structs.RoutePrefix{
		Prefix: "/api/users",
		SubRoutes: []structs.Route{
			{
				"Create a new user",
				"POST",
				"",
				server.CreateUser,
				true,
				[]gin.HandlerFunc{},
			},
			{
				"Get user data by ID",
				"GET",
				"/:id",
				server.GetUser,
				false,
				[]gin.HandlerFunc{},
			},
		},
	}
}
