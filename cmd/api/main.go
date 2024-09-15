package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/oneweerachai/service1/internal/api"
	"github.com/oneweerachai/service1/internal/db"
	"github.com/oneweerachai/service1/internal/external"
	log "github.com/oneweerachai/shared-logger/v2"
	"go.uber.org/zap"
)

func main() {
	logger, err := log.NewLogger()

	defer logger.Sync()

	if err != nil {
		panic(err)
	}
	fmt.Println("Hello")

	logger.Info("Hello, shared logger")
	logger.Debug("h debug")

	dbConn, err := db.ConnectDB()
	if err != nil {
		logger.Fatal("Failed to connect to the database", zap.Error(err))
	}

	defer dbConn.Close()

	// Initialize External API Client
	externalClient := external.NewClient("https://api.external.com", "API_KEY")

	router := gin.Default()

	api.RegisterRoutes(router, dbConn, externalClient, logger)

	if err := router.Run(":8080"); err != nil {
		logger.Fatal("Failed to run server", zap.Error(err))
	}
}
