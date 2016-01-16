package pkg

import "fmt"

type Square struct {
Length, Width int
Name          string
}

func (s Square) display() {
fmt.Printf("Length = %d, Width=%d, name = %s", s.Length, s.Width, s.Name)
return
}

type Circle struct {
	Length, Width int
	Name          string
}

func (c Circle) display() {
	fmt.Printf("Length = %d, Width=%d, name = %s", c.Length, c.Width, c.Name)
	return
}