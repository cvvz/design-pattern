package factory

import (
	"testing"
)

func TestNewPrinter(t *testing.T) {
	printer := NewIntPrinter(1, 2, 3, 4, 5)
	printer.Print()

	printer = NewCharPrinter('a', 'b', '我', '和', '你')
	printer.Print()
}
