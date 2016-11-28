package main 

import "fmt"

func main() {
	
	listNums := [] int {1, 2, 3, 4, 5}
	fmt.Println("Sum: ", addThemUp(listNums))

	// Funcs that return more than one value
	num1, num2 := some2Values(5)
	fmt.Println("Two values: ", num1, num2)

	// Undefined number args on a func
	fmt.Println(subtractThem(1, 2, 3, 4, 5))

	// Funcs on closure
	num3 := 3

	doubleNum := func() int{
		num3 *= 2
		return num3
	}
	
	fmt.Println(doubleNum())

	// Factorial
	fmt.Println(factorial(4))

	// Defer
	defer printTwo()
	printOne()

	// Defer an error
	fmt.Println(safeDiv(3,0))
	fmt.Println(safeDiv(3,2))
}

func addThemUp(numbers []int) int{
	
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum

}

func some2Values(number int) (int, int){
	
	return number + 1, number + 2

}

func subtractThem(args ...int) int{
	
	finalValue := 0

	for _, value := range args {
		finalValue -= value
	}

	return finalValue
}

func factorial(num int) int{
	if num == 0 {
		return 1
	}

	return num * factorial(num - 1)
}

func printOne(){
	fmt.Println((1))
}

func printTwo(){
	fmt.Println((2))
}

func safeDiv(num1, num2 int) int{
	
	defer func(){
		fmt.Println(recover())
	}()

	solution := num1 / num2
	return solution
}