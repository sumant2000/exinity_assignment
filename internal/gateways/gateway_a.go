package gateways

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sony/gobreaker"
	"net/http"
	"time"

	"Exinity/internal/utils"
)

// Define GatewayA struct
type GatewayA struct {
	URL            string
	circuitBreaker *gobreaker.CircuitBreaker
}

// Define PaymentRequest struct
type PaymentRequest struct {
	Amount    float64 `json:"amount"`
	AccountID string  `json:"accountID"`
}

// Define PaymentResponse struct
type PaymentResponse struct {
	TransactionID string `json:"transactionID"`
	Status        string `json:"status"`
}

// NewGatewayA creates a new instance of GatewayA
func NewGatewayA(url string) *GatewayA {
	return &GatewayA{
		URL: url,
	}
}

func (g *GatewayA) Deposit(amount float64, accountID string) (string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second, // Timeout after 10 seconds
	}

	// Retry logic wrapped inside the circuit breaker
	result, err := g.circuitBreaker.Execute(func() (interface{}, error) {
		return utils.Retry(3, 2*time.Second, func() (string, error) {
			// Create payment request
			request := PaymentRequest{Amount: amount, AccountID: accountID}
			reqBody, _ := json.Marshal(request)

			// Create a new HTTP request
			req, err := http.NewRequest("POST", fmt.Sprintf("%s/deposit", g.URL), bytes.NewBuffer(reqBody))
			if err != nil {
				return "", err
			}
			req.Header.Set("Content-Type", "application/json")

			// Execute the HTTP request
			resp, err := client.Do(req)
			if err != nil {
				return "", fmt.Errorf("error contacting gateway A: %v", err)
			}
			defer resp.Body.Close()

			// Parse the response
			var paymentResponse PaymentResponse
			if err := json.NewDecoder(resp.Body).Decode(&paymentResponse); err != nil {
				return "", fmt.Errorf("error parsing response from gateway A: %v", err)
			}

			return paymentResponse.TransactionID, nil
		})
	})

	if err != nil {
		return "", err
	}

	return result.(string), nil
}
