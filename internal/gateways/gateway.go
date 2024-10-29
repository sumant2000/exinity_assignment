package gateways

type PaymentGateway interface {
	Deposit(amount float64, accountID string) (string, error)
	Withdraw(amount float64, accountID string) (string, error)
}
