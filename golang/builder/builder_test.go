package builder

import "testing"

func TestBuilder(t *testing.T) {
	p := NewPrinterBuilder().
		WithInt(97).
		WithChar('a').
		Build()

	if err := p.Err(); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	p.Print()
}
