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
	a[3] = 100
	fmt.Println("a add a[3]", a)
	fmt.Println("set :", a)
	fmt.Println("get:", a[3])
	fmt.Println("len: ", len(a))

	println("-------------------------------------------------------")
	d := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("dcl:", d)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func sampleArrayString() {
	var aString [20]string
	fmt.Println("Empty String a", aString)

	bString := [3]string{}
	fmt.Println("String B", bString)
}

func main() {
	fmt.Println("Hello go day3 Array")
	fmt.Println("--------------------------------------------------")
	sampleArrayInt()
	fmt.Println("--------------------------------------------------")
	sampleArrayString()
}
