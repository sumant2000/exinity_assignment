package main

import (
	_ "Exinity/docs"
	"Exinity/internal/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func main() {
	r := gin.Default()

	// Swagger UI route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register routes (only once for each)
	r.GET("/health", handlers.HealthCheck)
	r.POST("/deposit", handlers.HandleDeposit)
	r.POST("/withdraw", handlers.HandleWithdraw)
	r.POST("/callback", handlers.HandleCallback)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
