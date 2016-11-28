package main

import ("fmt"
"strings"
"sort")
// "os"
// "log"
// "io/ioutil"
// "strconv")

func main() {
	
	sampString := "Hello World"

	fmt.Println(strings.Contains(sampString, "lo"))
	fmt.Println(strings.Index(sampString, "lo"))
	fmt.Println(strings.Count(sampString, "l"))
	fmt.Println(strings.Replace(sampString, "l", "x", 3))
	fmt.Println(strings.Replace(sampString, "l", "x", 2))


	csvString := "1,2,3,4,5,6"

	fmt.Println(strings.Split(csvString, ","))

	listOfLetters := [] string {"c", "a", "g"}

	sort.Strings(listOfLetters)

	fmt.Println("Letters: ", listOfLetters)


	listOfNums := strings.Join([] string {"3", "2", "1"}, ", ")
	fmt.Println("Numbers: ", listOfNums)
}