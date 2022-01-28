package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jacksontong/aoc-2017-day-1-go/captcha"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("please specify input.")
	}

	// extract the input to slice
	digits := transformInput(os.Args[1])

	// invoke SumNext
	answer := captcha.SumNext(digits...)

	// print result
	fmt.Printf("Answer: %d\n", answer)
}

// Take a string of a sequence of digits
// and return a slice of int
func transformInput(input string) []int {
	out := make([]int, len(input))

	for i, letter := range input {
		d, err := strconv.Atoi(string(letter))
		if err != nil {
			out[i] = 0
		} else {
			out[i] = d
		}
	}

	return out
}
