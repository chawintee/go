package main

import "fmt"

func sampleArrayInt() {
	// 	var <ชื่อของอาเรย์> [<ขนาดของอาเรย์>]<ประเภทของอาเรย์>  #1
	//     <ชื่อของอาเรย์> := [<ขนาดของอาเรย์>]<ประเภทของอาเรย์>{<กำหนดค่า>}    #2
	// var <ชื่อของอาเรย์> [<ขนาดของอาเรย์>][<ขนาดของอาเรย์>]<ประเภทของอาเรย์>     #3

	// #1
	var a [5]int
	fmt.Println("emp [a]:", a)

	// #2
	b := [8]int{2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("emp [b]:", b)

	// #3
	var c [2][3]int
	fmt.Println("2d ArrayC[2][3]", c)

	// note : if var not input value OK

}

func main() {
	fmt.Println("Hello go day3 Array")
	sampleArray()
}
