package main 

import "fmt"

func main() {

	// Slices are like arrays but without fixed slots
	numSlice := []int {9, 7, 6, 4, 3, 1, 5, 2, 8}

	// Creating slice from a nother slice
	numSlice2 := numSlice[3:5] // {4, 3}

	fmt.Println("numSlice2[0] = ", numSlice2[0])

	fmt.Println("numSlice[:2] = ", numSlice[:2])

	fmt.Println("numSlice[2:5] = ", numSlice[2:5])
	
	// Slice without pre-defined values
	// 				type,  many max
	numSlice3 := make([]int, 6, 10)

	copy(numSlice3, numSlice)

	for _, value := range numSlice3 {
		fmt.Println(value)
	}

	numSlice3 = append(numSlice3, 200, -1, 320)

	fmt.Println(numSlice3[6])
	fmt.Println(numSlice3[7])
	fmt.Println(numSlice3[8])
}