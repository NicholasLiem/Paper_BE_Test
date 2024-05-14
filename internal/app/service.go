package app

import "github.com/NicholasLiem/Paper_BE_Test/internal/service"

type MicroserviceServer struct {
	userService        service.UserService
	walletService      service.WalletService
	transactionService service.TransactionService
}

func NewMicroservice(
	userService service.UserService,
	walletService service.WalletService,
	transactionService service.TransactionService,
) *MicroserviceServer {
	return &MicroserviceServer{
		userService:        userService,
		walletService:      walletService,
		transactionService: transactionService,
	}
}
