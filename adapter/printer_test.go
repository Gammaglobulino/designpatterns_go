package adapter

import (
	"../adapter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintAdapterLegacyPrinter(t *testing.T) {
	msg := "Hello world"
	adapter := adapter.PrintAdapter{OldPrinter: &adapter.MyLegacyPrinter{}, Msg: msg}
	assert.EqualValues(t, adapter.PrintStored(), "Legacy Printer: Adapter: Hello world\n")
}

func TestPrintAdapterActualPrinter(t *testing.T) {
	msg := "Hello world"
	adapter := adapter.PrintAdapter{OldPrinter: nil, Msg: msg}
	assert.EqualValues(t, adapter.PrintStored(), "Hello world")
}
