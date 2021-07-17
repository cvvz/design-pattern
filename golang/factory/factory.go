package factory

import "fmt"

type Printer interface {
	Print()
}

type intPrinter struct {
	data []int
}

func (ip *intPrinter) Print() {
	for _, v := range ip.data {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

type charPrinter struct {
	data string
}

func (ap *charPrinter) Print() {
	for _, v := range ap.data {
		fmt.Printf("%c ", v)
	}
	fmt.Println()
}

func NewIntPrinter(data ...int) Printer {
	return &intPrinter{
		data: data,
	}
}

func NewCharPrinter(str ...rune) Printer {
	return &charPrinter{
		data: string(str),
	}
}
