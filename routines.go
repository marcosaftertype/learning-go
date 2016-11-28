package main

import ("fmt")

func main() {
	
	// Synchronous
	// thread1()
	// thread2()

	// Async
	go thread1() // Doesn't finishes the 20 because thread2 finishes
	thread2()
}

func thread1() {
	for i := 0; i < 20; i++ {
		fmt.Println("From thread1: ", i)
	}
}

func thread2() {
	for i := 0; i < 10; i++ {
		fmt.Println("From thread2: ", i)
	}
}