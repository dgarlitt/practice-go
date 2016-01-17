package shapes

import "fmt"

// Rectangle shape definition
type Rectangle struct {
	Length, Width int
	Name          string
}

func (r Rectangle) display() {
	fmt.Printf("Length = %d, Width=%d, Name = %s", r.Length, r.Width, r.Name)
	return
}

// Square shape definition
type Square struct {
	Width int
	Name  string
}

func (s Square) display() {
	fmt.Printf("Width=%d, Name = %s", s.Width, s.Name)
	return
}

// Circle shape definition
type Circle struct {
	Radius int
	Name   string
}

func (c Circle) display() {
	fmt.Printf("Radius=%d, Name = %s", c.Radius, c.Name)
	return
}
