package main

import "fmt"

type Rect struct {
	Width, Height int
}

func (r Rect) Area() int {
	return r.Width * r.Height
}

func (r Rect) DoubleSize()  {
	r.Width *= 2
	r.Height *= 2
	fmt.Println("DoubleSize()", r)
}

func (r Rect) String() string {
	return fmt.Sprintf("Rect{Width: %d, Height: %d}", r.Width, r.Height)
}

func main() {
	r := Rect{10, 5}
	fmt.Printf("Area of rectangle: %d\n", r.Area())
	r.DoubleSize()
}