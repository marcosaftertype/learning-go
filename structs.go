package main 

import "fmt"

func main() {
	
	rect1 := Rectangle {leftX: 0, topY: 50, height: 10, width: 12.5}

	// Same as
	// rect1 := Rectangle {0, 50, 10, 12.5}

	fmt.Println("Rectangle width is: ", rect1.width)
	fmt.Println("Rectangle area: ", rect1.area())
}

type Rectangle struct {
	leftX float64
	topY float64
	height float64
	width float64
}

// Define area method for rectangle struct
func (rect *Rectangle) area() float64{
	
	return rect.width * rect.height

}