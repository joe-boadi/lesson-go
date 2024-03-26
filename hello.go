package main

// import: Package quote collects pithy sayings.
import (
	"fmt"

	"rsc.io/quote"
)

// Variable Shadowing.
func example() {
	x := 1
	fmt.Println(x)

	if true {
		fmt.Println(x)
		x := 2
		fmt.Println(x)
	}

	fmt.Println(x)
}

// Array & for loop.
func array_() {
	//var arrayList []string ,slice
	arrayList := []string{"Books", "Pens", "Pencils"}
	for _, element := range arrayList {
		fmt.Println(element)
	}
}

func myMap() {
	// myMap := make(map[string]int)
	var myMap = map[string]int{"Joe": 17, "Eve": 16}
	fmt.Printf("Joe has %v cards", myMap["Joe"])
}
func main() {
	//returns statements from "rsc".
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
	fmt.Println(quote.Glass())
	example()
	array_()
	myMap()
}
