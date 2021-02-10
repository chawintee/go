package main

import (
	"fmt"
	"math"
)

func ifElse() {
	if 7%2 == 0 {
		fmt.Println("7 is even number")
	} else {
		fmt.Println("7 is odd number")
	}
}

func ifElsewithVariable() {
	i := 30
	if i%2 == 0 {
		fmt.Println(i, "is even number")
	} else {
		fmt.Println(i, "is odd number")
	}
}

func ifCondition() {
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
}

func ifElseIfCondition() {
	if num := 3; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

func main() {
	fmt.Println("Hello Day3 If else")
	var powTest float64
	powTest = math.Pow(2, 5)
	fmt.Println(powTest)
	fmt.Println("-----------------------------------------------------------------------")
	ifElse()
	fmt.Println("-----------------------------------------------------------------------")
	ifElsewithVariable()
	fmt.Println("-----------------------------------------------------------------------")
	ifCondition()
	fmt.Println("-----------------------------------------------------------------------")
	ifElseIfCondition()
	fmt.Println("-----------------------------------------------------------------------")
}
