package handlers

import (
	"Exinity/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/deposit", HandleDeposit)
	r.POST("/withdraw", HandleWithdraw)
	r.POST("/callback", HandleCallback)
}

// Initialize the transactionService
var transactionService = services.NewTransactionService() // Assuming you have this method in your services package

// @Summary Deposit funds
// @Description Deposit funds into an account
// @Accept  json
// @Produce  json
// @Param   amount    body  float64  true  "Amount to deposit"
// @Param   accountID body  string   true  "Account ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /deposit [post]
// HandleDeposit handles deposit requests
func HandleDeposit(c *gin.Context) {
	var request struct {
		Amount    float64 `json:"amount"`
		AccountID string  `json:"accountID"`
		Gateway   string  `json:"gateway"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	transactionID, err := transactionService.ProcessDeposit(request.Amount, request.AccountID, request.Gateway)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactionID": transactionID})
}

// HandleWithdraw handles withdraw requests
func HandleWithdraw(c *gin.Context) {
	var request struct {
		Amount    float64 `json:"amount"`
		AccountID string  `json:"accountID"`
		Gateway   string  `json:"gateway"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	transactionID, err := transactionService.ProcessWithdraw(request.Amount, request.AccountID, request.Gateway)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactionID": transactionID})
}

func HandleCallback(c *gin.Context) {
	// Logic for handling gateway callbacks
}
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Server is running",
	})
}
