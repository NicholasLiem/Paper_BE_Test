package app

import (
	"github.com/NicholasLiem/Paper_BE_Test/utils"
	"github.com/NicholasLiem/Paper_BE_Test/utils/http"
	"github.com/gin-gonic/gin"
	netHttp "net/http"
)

func (m *MicroserviceServer) GetTransaction(c *gin.Context) {
	transactionID := c.Param("id")

	parsedTransactionID, err := utils.ParseStrToUint(transactionID)
	if err != nil {
		http.ErrorResponse(c.Writer, netHttp.StatusBadRequest, "Invalid transaction id")
		return
	}

	result, httpErr := m.transactionService.GetTransaction(*parsedTransactionID)
	if httpErr != nil {
		http.ErrorResponse(c.Writer, httpErr.StatusCode, httpErr.Message)
		return
	}

	http.SuccessResponse(c.Writer, netHttp.StatusOK, "Transaction retrieved successfully", result)
}

func (m *MicroserviceServer) GetTransactionsByWalletID(c *gin.Context) {
	walletID := c.Param("wallet_id")

	parsedWalletID, err := utils.ParseStrToUint(walletID)
	if err != nil {
		http.ErrorResponse(c.Writer, netHttp.StatusBadRequest, "Invalid wallet id")
		return
	}

	result, httpErr := m.transactionService.GetTransactionsByWalletID(*parsedWalletID)
	if httpErr != nil {
		http.ErrorResponse(c.Writer, httpErr.StatusCode, httpErr.Message)
		return
	}

	http.SuccessResponse(c.Writer, netHttp.StatusOK, "Transactions retrieved successfully", result)
}
