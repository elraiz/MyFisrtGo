package main

import (
	"fmt"
)

func main() {
	var width = 10
	var height = 512
	var number = 7

	calculate(width, height)
	isNumberPrime(number)
}

func calculate(width int, height int) {
	for i := 0; i < width; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	if width > 0 && height > 0 {
		var area = (width * height) / 2

		fmt.Println("The square are of the triangle is : ", area)
	} else {
		fmt.Println("Invalid width or height")
	}
}

func isNumberPrime(number int) {
	var divider int
	if number < 3 {
		fmt.Println("number ", number, " is a prime number")
	} else {
		for i := 2; i <= number; i++ {
			divider = i
			if (number % i) == 0 {
				break
			}
		}
		if divider == number {
			fmt.Println("number ", number, " is a prime number")
		} else {
			fmt.Println("number ", number, " is not a prime number")
		}
	}
}
