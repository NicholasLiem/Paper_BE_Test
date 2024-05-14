package app

import (
	"github.com/NicholasLiem/Paper_BE_Test/internal/datastruct"
	"github.com/NicholasLiem/Paper_BE_Test/utils"
	"github.com/NicholasLiem/Paper_BE_Test/utils/http"
	"github.com/gin-gonic/gin"
	netHttp "net/http"
)

func (m *MicroserviceServer) CreateUser(c *gin.Context) {
	var user datastruct.User
	if err := c.ShouldBindJSON(&user); err != nil {
		http.ErrorResponse(c.Writer, netHttp.StatusBadRequest, "Invalid request payload")
		return
	}

	result, httpErr := m.userService.CreateUser(user)
	if httpErr != nil {
		http.ErrorResponse(c.Writer, httpErr.StatusCode, httpErr.Message)
		return
	}

	http.SuccessResponse(c.Writer, netHttp.StatusCreated, "User created successfully", result)
}

func (m *MicroserviceServer) GetUser(c *gin.Context) {
	userID := c.Param("id")

	parsedUserID, err := utils.ParseStrToUint(userID)
	if err != nil {
		http.ErrorResponse(c.Writer, netHttp.StatusBadRequest, "Invalid user id")
		return
	}

	result, httpErr := m.userService.GetUser(*parsedUserID)
	if httpErr != nil {
		http.ErrorResponse(c.Writer, httpErr.StatusCode, httpErr.Message)
		return
	}

	http.SuccessResponse(c.Writer, netHttp.StatusOK, "User retrieved successfully", result)
}
