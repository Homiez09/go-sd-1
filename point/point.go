package point

import "fmt"

type Point struct {
	x int
	y int
}

func NewPoint(x int, y int) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) Print() {
	fmt.Println(p.x, p.y)
}
