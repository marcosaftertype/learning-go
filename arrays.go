package main 

import "fmt"

func main() {
	
	// First approach
	var favNums[5] float64

	favNums[0] = 163
	favNums[1] = 789
	favNums[2] = 1.0325
	favNums[3] = 2.567
	favNums[4] = 43216579

	fmt.Println(favNums[3])

	// Second approach
	favNums2 := [5]float64 {1,2,3,4,5}

	for i, value := range favNums2  {	
		fmt.Println(value, i)
	}

	// If no need for index
	for _, value := range favNums2  {	
		fmt.Println(value)
	}

}