package factory

import (
	"../factory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPaymentMethodCash(t *testing.T) {
	payment, err := factory.GetPaymentMethod(factory.Cash)
	if err != nil {
		t.Fatal("A payment method of type Cash must exist")
	}
	msg := payment.Pay(10.30)
	assert.EqualValues(t, "10.30 paid using cash", msg)
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodDebit(t *testing.T) {
	payment, err := factory.GetPaymentMethod(factory.DebitCard)
	if err != nil {
		t.Fatal("A payment method of type Debit must exist")
	}
	msg := payment.Pay(10.30)
	assert.EqualValues(t, "10.30 paid using Debit Card", msg)
	t.Log("LOG:", msg)
}
func TestGetPaymentMethodNetDefined(t *testing.T) {
	_, err := factory.GetPaymentMethod(20)
	if err == nil {
		t.Error("A payment method with id 20 must exist")
	}
	t.Log("LOG:", err)
}
