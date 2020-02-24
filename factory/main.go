package factory

import (
	"errors"
	"fmt"
)

// delegating the creation of different payments methods

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash = iota
	DebitCard
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return &CashPM{}, nil
	case DebitCard:
		return &DebitCardPM{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized\n", m))
	}
}

type CashPM struct{}
type DebitCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash", amount)
}
func (d *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using Debit Card", amount)
}
