package decorator

import "testing"

func TestNewPrinter(t *testing.T) {
	p := NewPrinter()
	p.Print()

	p1 := NewPrinter(WithInt(1, 2, 3, 4, 5))
	p1.Print()

	p2 := NewPrinter(WithChar('我', '和', '你', '!'))
	p2.Print()
}
