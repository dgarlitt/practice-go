package shapes

func ExampleRectangle() {
	rectangle := Rectangle{3, 4, "myRectangle"}
	rectangle.display()
	// Output: Length = 3, Width=4, Name = myRectangle
}

func ExampleSquare() {
	rectangle := Square{4, "myRectangle"}
	rectangle.display()
	// Output: Width=4, Name = myRectangle
}

func ExampleCircle() {
	circle := Circle{2, "myCircle"}
	circle.display()
	// Output: Radius=2, Name = myCircle
}
