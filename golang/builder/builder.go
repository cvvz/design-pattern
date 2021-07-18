package builder

import "fmt"

type IPrinter interface {
	Err() error
	Print()
}

type Printer struct {
	// 🌟🌟注意这个把error作为结构体中的一个成员变量的技巧，可以避免Error Check Hell
	err error

	i int
	c rune
}

func (p *Printer) Print() {
	fmt.Printf("%d %c\n", p.i, p.c)
}

func (p *Printer) Err() error {
	return p.err
}

type PrinterBuilder struct {
	Printer
}

func NewPrinterBuilder() *PrinterBuilder {
	return &PrinterBuilder{}
}

func (pb *PrinterBuilder) WithInt(i int) *PrinterBuilder {
	pb.Printer.i = i
	return pb
}

func (pb *PrinterBuilder) WithChar(c rune) *PrinterBuilder {
	pb.Printer.c = c
	return pb
}

func (pb *PrinterBuilder) Build() IPrinter {
	if int(pb.Printer.c) == pb.Printer.i {
		pb.Printer.err = fmt.Errorf("the value of c is equal to i %d", pb.Printer.c)
	}
	return &pb.Printer
}
