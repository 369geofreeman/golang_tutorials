package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

// package main

// import "fmt"

// type Counter struct {
// 	n int
// }

// func (c Counter) Count() {
// 	for c.n != 0 {
// 		fmt.Println(c.n)
// 		c.n -= 1
// 	}
// 	fmt.Println("Go!")
// }

// func main() {
// 	c := Counter{n: 3}
// 	c.Count()
// }
