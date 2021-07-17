package decorator

import "fmt"

type Printer interface {
	Print()
}

type printer struct {
	i []int
	c []rune
}

type OptionFunc func(p *printer) *printer

func NewPrinter(options ...OptionFunc) Printer {
	p := &printer{}

	for _, opt := range options {
		p = opt(p)
	}

	return p
}

func WithInt(data ...int) OptionFunc {
	return func(p *printer) *printer {
		// copy(p.i, data) 由于p.i长度为0，这个地方不能直接copy
		p.i = data
		return p
	}
}

func WithChar(data ...rune) OptionFunc {
	return func(p *printer) *printer {
		p.c = make([]rune, len(data))
		copy(p.c, data)
		return p
	}
}

func (p *printer) Print() {
	fmt.Printf("printer: \n")
	if len(p.c) != 0 {
		for _, c := range p.c {
			fmt.Printf("%c ", c)
		}
		fmt.Println()
	}
	if len(p.i) != 0 {
		for _, i := range p.i {
			fmt.Printf("%d ", i)
		}
		fmt.Println()
	}
}
