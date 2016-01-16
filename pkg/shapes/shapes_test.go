package shapes

func ExampleRectangle() {
	rectangle := Rectangle{3, 4, "myRectangle"}
	rectangle.display()
	// Output: Length = 3, Width=4, name = myRectangle
}

func ExampleCircle() {
	circle := Circle{3, 2, "myCircle"}
	circle.display()
}
