package main

import "github.com/Homiez09/go-sd-1.git/printer"

func Print(pt printer.Printer) {
	pt.Print()
}

func main() {
	// inkjet := printer.NewInkjetPrinter("InkJet") // return เป็น pointer อยู่แล้ว
	laser := printer.NewLaserPrinter("Laser") // return เป็น pointer อยู่แล้ว
	Print(laser)
}
