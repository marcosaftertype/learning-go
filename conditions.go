package main

import "fmt"

func main() {
	
	yourAge := 21

	// If
	fmt.Println("If statement")

	if yourAge >= 21 {
		fmt.Println("Ok, come in")
	} else if yourAge >= 50 {
		fmt.Println("Have you tried our pancakes?")
	} else {
		fmt.Println("You are still a kiddo")
	}

	// Switch
	fmt.Println("Switch statement")

	anotherAge := 50

	switch anotherAge {
		case 21 : fmt.Println("Ok, come in")
		case 50 : fmt.Println("Pancakes")
		default : fmt.Println("You are still a kiddo")
	}

}