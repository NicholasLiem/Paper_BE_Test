package app

import (
	"github.com/NicholasLiem/Paper_BE_Test/utils"
	"github.com/NicholasLiem/Paper_BE_Test/utils/http"
	"github.com/gin-gonic/gin"
	netHttp "net/http"
)

func (m *MicroserviceServer) Topup(c *gin.Context) {
	var req struct {
		UserID uint    `json:"user_id"`
		Amount float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		http.ErrorResponse(c.Writer, netHttp.StatusBadRequest, "Invalid request payload")
		return
	}

	result, httpErr := m.walletService.Topup(req.UserID, req.Amount)
	if httpErr != nil {
		http.ErrorResponse(c.Writer, httpErr.StatusCode, httpErr.Message)
		return
	}

	http.SuccessResponse(c.Writer, netHttp.StatusOK, "Top up successful", result)
}

func (m *MicroserviceServer) Withdraw(c *gin.Context) {
	var req struct {
		UserID uint    `json:"user_id"`
		Amount float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		http.ErrorResponse(c.Writer, netHttp.StatusBadRequest, "Invalid request payload")
		return
	}

	result, httpErr := m.walletService.Withdraw(req.UserID, req.Amount)
	if httpErr != nil {
		http.ErrorResponse(c.Writer, httpErr.StatusCode, httpErr.Message)
		return
	}

	http.SuccessResponse(c.Writer, netHttp.StatusOK, "Withdrawal successful", result)
}

func (m *MicroserviceServer) GetBalance(c *gin.Context) {
	userID := c.Param("user_id")

	parsedUserID, err := utils.ParseStrToUint(userID)
	if err != nil {
		http.ErrorResponse(c.Writer, netHttp.StatusBadRequest, "Invalid user id")
		return
	}

	result, httpErr := m.walletService.GetBalance(*parsedUserID)
	if httpErr != nil {
		http.ErrorResponse(c.Writer, httpErr.StatusCode, httpErr.Message)
		return
	}

	http.SuccessResponse(c.Writer, netHttp.StatusOK, "Balance retrieved successfully", gin.H{"balance": result})
}
