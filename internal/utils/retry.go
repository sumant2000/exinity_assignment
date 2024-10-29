package utils

import (
	"fmt"
	"time"
)

// Retry function with configurable attempts and sleep duration
func Retry(attempts int, sleep time.Duration, fn func() (string, error)) (string, error) {
	for i := 0; i < attempts; i++ {
		result, err := fn()
		if err == nil {
			return result, nil
		}
		time.Sleep(sleep)
	}
	return "", fmt.Errorf("max retries reached")
}
