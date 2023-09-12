package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Make the computer guess a number, we will tell the computer if it is too high/low

	scanner := bufio.NewScanner(os.Stdin)

	low := 1
	high := 10

	fmt.Println("Please think of a number between", low, "and", high)
	fmt.Println("Press enter when ready")

	scanner.Scan()

	// n = 1 to 100 O(n) Big O notation

	for {

		// Binary Search Strategy

		guess := (low + high) / 2
		fmt.Println("I guess the number is", guess)
		fmt.Println("Is that:")
		fmt.Println("(a) Calm down Cheech, you're way too high?")
		fmt.Println("(b) Way too low, you kind of suck at this?")
		fmt.Println("(c) Am I correct?")
		scanner.Scan()
		response := scanner.Text()

		if response == "a" {
			high = guess - 1
		} else if response == "b" {
			low = guess + 1
		} else if response == "c" {
			fmt.Println("Winner, winner, chicken dinner!")
			break
		} else {
			fmt.Println("Invalid response, try again")
		}
	}
}
