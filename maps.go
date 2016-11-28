package main 

import "fmt"

func main() {

	fruitsBasket := make(map[string] int8)
	
	fruitsBasket["apples"] = 42
	fruitsBasket["oranges"] = 28
	fruitsBasket["grapes"] = 130

	fmt.Println(fruitsBasket["apples"])
	fmt.Println(fruitsBasket["oranges"])
	
	fmt.Println(len(fruitsBasket))

	delete(fruitsBasket, "apples")
	
	fmt.Println(len(fruitsBasket))
}