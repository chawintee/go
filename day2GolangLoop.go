package main

import "fmt"

func init1() {
	fmt.Println("int... Golang Day2 Loop...")
}

func forLoop() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

func forLoopBreak() {
	for j := 0; j <= 100; j++ {
		if j == 5 {
			continue
		} else if j == 10 {
			break
		}
		fmt.Println(j)
	}
}

func main() {
	init1()
	fmt.Println("Hello GOLANG Loop...")
	// forLoop()
	forLoopBreak()
}
