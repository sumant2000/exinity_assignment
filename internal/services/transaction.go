package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Example function
func (s *TransactionService) ProcessTransaction() {
	// Your logic
}

// TransactionService handles transactions
type TransactionService struct {
	// Add fields like gateway or other dependencies here
}

// NewTransactionService creates a new instance of TransactionService
func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

// Define the ProcessWithdraw method for TransactionService
func (s *TransactionService) ProcessWithdraw(amount float64, accountID, gatewayType string) (string, error) {
	// Add logic for processing a withdrawal based on the gateway type
	if gatewayType == "A" {
		// Process withdrawal with Gateway A
		return "withdrawal_id_from_gateway_a", nil
	} else if gatewayType == "B" {
		// Process withdrawal with Gateway B
		return "withdrawal_id_from_gateway_b", nil
	}
	return "", fmt.Errorf("unknown gateway type")
}

// Define the ProcessDeposit method for TransactionService
func (s *TransactionService) ProcessDeposit(amount float64, accountID, gatewayType string) (string, error) {
	// Add logic for processing a deposit based on the gateway type
	// For example, call the appropriate gateway's deposit method (Gateway A or B)

	// Mock implementation:
	if gatewayType == "A" {
		// Process with Gateway A
		return "transaction_id_from_gateway_a", nil
	} else if gatewayType == "B" {
		// Process with Gateway B
		return "transaction_id_from_gateway_b", nil
	}
	return "", fmt.Errorf("unknown gateway type")
}

var transactionService *TransactionService

// Initialize the service in main.go (not shown)

func HandleDeposit(c *gin.Context) {
	var json struct {
		Amount    float64 `json:"amount"`
		AccountID string  `json:"accountID"`
		Gateway   string  `json:"gateway"`
	}

	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	transactionID, err := transactionService.ProcessDeposit(json.Amount, json.AccountID, json.Gateway)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactionID": transactionID})
}

func HandleWithdraw(c *gin.Context) {
	var json struct {
		Amount    float64 `json:"amount"`
		AccountID string  `json:"accountID"`
		Gateway   string  `json:"gateway"`
	}

	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	transactionID, err := transactionService.ProcessWithdraw(json.Amount, json.AccountID, json.Gateway)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactionID": transactionID})
}
