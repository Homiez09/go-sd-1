package printer

import "fmt"

type LaserPrinter struct {
	Name string
}

func NewLaserPrinter(name string) Printer {
	return &LaserPrinter{Name: name}
}

func (i *LaserPrinter) Print() {
	// การทำงานของ Laser
	fmt.Println("Prepare ink for LASER")

	// ผลลัพธ์
	fmt.Println("PRINT")
}
