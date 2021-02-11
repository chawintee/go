package main

import "fmt"

func sampleSlices() {
	// 	<ชื่อของSlice> := make([]<ประเภทของSlice>, <ขนาดเริ่มต้น>)
	// <ชื่อของSlice> := []string{<ค่าเริ่มต้น>, <ค่าเริ่มต้น>}
	s := make([]string, 3)
	fmt.Println("emptySliceName:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set : ", s)
	fmt.Println("get : ", s[1])

	fmt.Println("lenOfS : ", len(s))

	s = append(s, "d")
	fmt.Println("append1 : ", s)
	fmt.Println("lenOfSAfterAppend1 : ", len(s))

	s = append(s, "e", "f")
	fmt.Println("append2 : ", s)
	fmt.Println("lenOfSAfterAppend2 : ", len(s))

}

func main() {
	fmt.Println("Hello Day3 Golang Slices")
	fmt.Println("----------------------------------------------------")
	sampleSlices()
}
