package main 

import "fmt"

func main() {
	
	x := 0

	// Doesn't change x
	changeXVal(x)
	fmt.Println("x = ", x)

	// Pointers do change the x
	changeXValNow(&x)
	fmt.Println("x = ", x)
}

func changeXVal(x int) {
	
	x = 2

}

func changeXValNow(x *int) {

	*x = 2

}