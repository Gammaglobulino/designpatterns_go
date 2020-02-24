package adapter

import "fmt"

type LegacyPrinter interface {
	Print(s string) string
}
type MyLegacyPrinter struct{}

func (p *MyLegacyPrinter) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy Printer: %s\n", s)
	return
}

type ModernPrinter interface {
	PrintStored() string
}
type PrintAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrintAdapter) PrintStored() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
	} else {
		newMsg = p.Msg
	}
	return
}
