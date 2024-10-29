package gateways

import (
	"Exinity/internal/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test ProcessDeposit for Gateway A
func TestProcessDepositGatewayA(t *testing.T) {
	service := services.NewTransactionService()

	transactionID, err := service.ProcessDeposit(100.0, "123456", "A")
	assert.NoError(t, err)
	assert.Equal(t, "transaction_id_a", transactionID)
}

// Test ProcessDeposit for Gateway B
func TestProcessDepositGatewayB(t *testing.T) {
	service := services.NewTransactionService()

	transactionID, err := service.ProcessDeposit(100.0, "123456", "B")
	assert.NoError(t, err)
	assert.Equal(t, "transaction_id_b", transactionID)
}

// Test ProcessDeposit for unknown gateway
func TestProcessDepositUnknownGateway(t *testing.T) {
	service := services.NewTransactionService()

	transactionID, err := service.ProcessDeposit(100.0, "123456", "C")
	assert.Error(t, err)
	assert.Empty(t, transactionID)
}

// Test ProcessWithdraw for Gateway A
func TestProcessWithdrawGatewayA(t *testing.T) {
	service := services.NewTransactionService()

	transactionID, err := service.ProcessWithdraw(50.0, "123456", "A")
	assert.NoError(t, err)
	assert.Equal(t, "withdrawal_id_a", transactionID)
}

// Test ProcessWithdraw for Gateway B
func TestProcessWithdrawGatewayB(t *testing.T) {
	service := services.NewTransactionService()

	transactionID, err := service.ProcessWithdraw(50.0, "123456", "B")
	assert.NoError(t, err)
	assert.Equal(t, "withdrawal_id_b", transactionID)
}

// Test ProcessWithdraw for unknown gateway
func TestProcessWithdrawUnknownGateway(t *testing.T) {
	service := services.NewTransactionService()

	transactionID, err := service.ProcessWithdraw(50.0, "123456", "C")
	assert.Error(t, err)
	assert.Empty(t, transactionID)
}
