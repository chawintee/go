package main

import (
	"fmt"
)

// switchCase type 1 is no Break different other language
func switchCasetype1() {
	switch i := 5; i {
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

func main() {
	fmt.Println("Hello Day3 Golang SwitchCase")
	switchCasetype1()

}
