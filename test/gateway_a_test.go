package gateways_test

import (
	"Exinity/internal/gateways" // Ensure you're importing the correct package
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test the Deposit function for GatewayA
func TestGatewayADepositSuccess(t *testing.T) {
	// Test logic
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"transactionID": "txn_123", "status": "success"}`))
	}))
	defer server.Close()

	// Call the exported NewGatewayA function
	gatewayA := gateways.NewGatewayA(server.URL)

	transactionID, err := gatewayA.Deposit(100.0, "123456")

	assert.NoError(t, err)
	assert.Equal(t, "txn_123", transactionID)
}

// Test the Deposit function with a failure (e.g., gateway is down)
func TestGatewayADepositFailure(t *testing.T) {
	// Simulate a failed HTTP server (unreachable)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	// Initialize Gateway A with the test server URL
	gatewayA := gateways.NewGatewayA(server.URL)

	// Perform a deposit operation
	_, err := gatewayA.Deposit(100.0, "123456")

	// Check that an error occurred
	assert.Error(t, err)
}
