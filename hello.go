package main

// import: Package quote collects pithy sayings.
import (
	"fmt"
	"time"

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
	// myMap := make(map[string]int) "or"
	var myMap = map[string]int{"Joe": 17, "Eve": 16}
	fmt.Printf("Joe has %v cards \n", myMap["Joe"])

	var cards, ok = myMap["Jones"]
	if ok {
		fmt.Printf("Card is %v", cards)
	} else {
		fmt.Println("Invalid name")
	}
}

func timeTaken(slice []int, n int) time.Duration {
	t0 := time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}

// Structs and Interfaces
type myBook struct {
	ID        uint
	Title     string
	Published uint
	writer
}

type writer struct {
	name string
}

// TODO: Go over interfaces!

// Pointers % memory addressing: tackling memory redundancy
func squared(something *[5]float64) [5]float64 {
	for i := range something {
		something[i] = something[i] * something[i]
	}
	fmt.Printf("Memory address of something: %p\n", something)
	return *something
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

	n := 10000
	testSlice := []int{}
	testSlice2 := make([]int, 0, n)
	fmt.Printf("\nTotal time with pre-allocation: %v \n", timeTaken(testSlice2, n))
	fmt.Printf("Total time without pre-allocation: %v \n", timeTaken(testSlice, n))

	// Structs and Interfaces
	var newBook myBook = myBook{001, "Golang 101", 2024, writer{"Joe"}}
	fmt.Printf("\n%v\n", newBook)

	// Pointers & address: squared()
	var something_again = [5]float64{2, 3, 4, 5, 6}
	var result [5]float64 = squared(&something_again)
	fmt.Printf("\nResults from slice/arr: %v\n", result)
	fmt.Printf("Memory address of something_again: %p\n", &something_again)
}
