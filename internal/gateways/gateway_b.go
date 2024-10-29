package gateways

import (
	"bytes"
	"fmt"
	"github.com/sony/gobreaker"
	"io/ioutil"
	"net/http"
	"time"

	"Exinity/internal/utils"
)

// Define GatewayB struct
type GatewayB struct {
	URL            string
	circuitBreaker *gobreaker.CircuitBreaker
}

// Implement other functions...

func (g *GatewayB) Deposit(amount float64, accountID string) (string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second, // Timeout after 10 seconds
	}

	// Retry logic wrapped inside the circuit breaker
	result, err := g.circuitBreaker.Execute(func() (interface{}, error) {
		return utils.Retry(3, 2*time.Second, func() (string, error) {
			// Create the SOAP request for deposit
			soapBody := fmt.Sprintf(`
				<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
					<soapenv:Body>
						<deposit>
							<amount>%f</amount>
							<accountID>%s</accountID>
						</deposit>
					</soapenv:Body>
				</soapenv:Envelope>`, amount, accountID)

			req, err := http.NewRequest("POST", g.URL, bytes.NewBuffer([]byte(soapBody)))
			if err != nil {
				return "", err
			}
			req.Header.Set("Content-Type", "text/xml")

			// Execute the HTTP request
			resp, err := client.Do(req)
			if err != nil {
				return "", fmt.Errorf("error contacting gateway B: %v", err)
			}
			defer resp.Body.Close()

			// Read the response body
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return "", fmt.Errorf("error reading response from gateway B: %v", err)
			}

			// Print or process the response body (mocking response for now)
			fmt.Println("SOAP Response Body:", string(body))

			return "transaction_b_123", nil
		})
	})

	if err != nil {
		return "", err
	}

	return result.(string), nil
}
