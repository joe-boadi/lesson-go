package main

// import: Package quote collects pithy sayings.
import (
	"fmt"

	"rsc.io/quote"
)

// Variable Shadowing
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

func main() {

	//returns statements from "rsc".
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
	fmt.Println(quote.Glass())

	example()
}
