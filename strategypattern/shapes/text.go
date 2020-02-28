package shapes

import (
	"../strategy"
)

type TextSquare struct {
	strategy.PrintOutput
}

func (tx *TextSquare) Print() error {
	tx.Writer.Write([]byte("Circle"))
	return nil
}
