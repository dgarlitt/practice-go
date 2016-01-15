package main

import "fmt"

// Rectangle shape definition
type Rectangle struct {
	Length, Width int
	Name          string
}

func (r Rectangle) display() {
	fmt.Printf("Length = %d, Width=%d, name = %s", r.Length, r.Width, r.Name)
	return
}
