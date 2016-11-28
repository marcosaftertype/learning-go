package main 

import "fmt"

func main(){

	// For loop
	i := 1 

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// One liner
	var count int = 5

	for j := 0; j < count; j++ {
		fmt.Println(j)
	}

}