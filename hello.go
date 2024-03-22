package main

// import: Package quote collects pithy sayings.
import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, world")
	//returns a Go proverb.
	fmt.Println(quote.Go())
	// Display: "Don't communicate by sharing memory, share memory by communicating"
}
