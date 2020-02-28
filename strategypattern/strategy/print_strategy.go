package strategy

import "io"

type PrintStrategy interface {
	Print() error
	SetLog(writer io.Writer)
	SetWriter(writer io.Writer)
}

type PrintOutput struct {
	Writer    io.Writer
	LogWriter io.Writer
}

func (d *PrintOutput) SetLog(w io.Writer) {
	d.LogWriter = w
}
func (d *PrintOutput) SetWriter(w io.Writer) {
	d.Writer = w
}
