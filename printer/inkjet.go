package printer

import "fmt"

type InkjetPrinter struct {
	Name string
}

// class InkjetPrinter implements Printer คล้ายๆ
func NewInkjetPrinter(name string) Printer {
	return &InkjetPrinter{Name: name}
}

func (i *InkjetPrinter) Print() { // ถ้าลบ func นี้จะ warning เพราะว่าไม่ implement interface
	// การทำงานของ Inkjet
	fmt.Println("Prepare ink for INKJET")
	fmt.Println("EITEIT")

	// ผลลัพธ์
	fmt.Println("PRINT")
}

func (i *InkjetPrinter) Clean() {}
