package main

import (
	"github.com/NicholasLiem/Paper_BE_Test/adapter"
	"github.com/NicholasLiem/Paper_BE_Test/internal/app"
	"github.com/NicholasLiem/Paper_BE_Test/internal/datastruct"
	"github.com/NicholasLiem/Paper_BE_Test/internal/repository"
	"github.com/NicholasLiem/Paper_BE_Test/internal/service"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func main() {
	/**
	Load env file
	*/
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	/**
	DB setup
	*/
	db := repository.SetupDB()

	/**
	Registering DAO's and Services
	*/
	dao := repository.NewDAO(db)

	userService := service.NewUserService(dao)
	transactionService := service.NewTransactionService(dao)
	walletService := service.NewWalletService(dao, transactionService)

	/**
	Registering Services to Server
	*/
	server := app.NewMicroservice(
		userService,
		walletService,
		transactionService,
	)

	/**
	DB Migration
	*/
	if err := datastruct.Migrate(db, &datastruct.User{}, &datastruct.Wallet{}, &datastruct.Transaction{}); err != nil {
		log.Fatalf("Error during migration: %v", err)
	}

	serverRouter := adapter.NewRouter(*server)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		ExposedHeaders:   []string{},
	})

	handler := c.Handler(serverRouter)

	port := os.Getenv("APP_PORT")
	log.Println("[Server] Running the server on port " + port)

	if os.Getenv("ENVIRONMENT") == "DEV" {
		log.Fatal(http.ListenAndServe("127.0.0.1:"+port, handler))
	} else {
		log.Fatal(http.ListenAndServe(":"+port, handler))
	}
}
