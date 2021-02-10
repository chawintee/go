package main

import (
	"fmt"
)

// switchCase type 1 is no Break different other language
func switchCasetype1() {
	i := 5
	fmt.Print("Write", i, "to")
	switch i {
	case 1:
		fmt.Println("Hello 1")
	case 2:
		fmt.Println("Hello 2")
	case 3:
		fmt.Println("Hello 3")
	case 4:
		fmt.Println("Hello 4")
	case 5:
		fmt.Println("Hello 5")
	case 6:
		fmt.Println("Hello 6")
	default:
		fmt.Println(i, "not number in case")
	}
}

func switchCaseType2() {
	j := 3
	switch j {
	case 5:
		fmt.Println("five not three")
		fallthrough
	case 4:
		fmt.Println("four not three")
		fallthrough
	case 3:
		fmt.Print("three ")
		fallthrough
	case 2:
		fmt.Print("two ")
		fallthrough
	case 1:
		fmt.Print("one ")
		fallthrough
	case 0:
		fmt.Print("zero ")
		fallthrough
	default:
		fmt.Print("Go!!!")
	}
}

func main() {
	fmt.Println("Hello Day3 Golang SwitchCase")
	switchCasetype1()
	switchCaseType2()

}
