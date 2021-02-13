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

	//append similar push
	s = append(s, "d")
	fmt.Println("append1 : ", s)
	fmt.Println("lenOfSAfterAppend1 : ", len(s))

	s = append(s, "e", "f")
	fmt.Println("append2 : ", s)
	fmt.Println("lenOfSAfterAppend2 : ", len(s))

	//copy ,len = lenght
	c := make([]string, len(s))
	fmt.Println("empty c : ", c)
	copy(c, s)
	fmt.Println("c copy from s : ", c)

	c = append(c, "g", "h")
	fmt.Println("c append : ", c)

	s[0] = "5"
	c[0] = "25"
	fmt.Println("c Change c[0]", c)
	fmt.Println("s change s[0]", s)
}

func sliceWithHightLow() {
	sl0 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("sl0 : ", sl0)

	sl0 = sl0[3:5]
	fmt.Println("sl0[3:5] : ", sl0)

	sl1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sl1 = sl1[:5]
	fmt.Println("sl1[:5] : ", sl1)

	fmt.Println("----------------------------------------------------")
	sl2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("sl2 : ", sl2)
	fmt.Printf("pointer of sl2[5] : %d , Address : %p\n", sl2[5], &sl2[5])
	fmt.Printf("pointer of sl2[6] : %d , Address : %p\n", sl2[6], &sl2[6])
	fmt.Printf("pointer of sl2[7] : %d , Address : %p\n", sl2[7], &sl2[7])
	fmt.Printf("pointer of sl2[8] : %d , Address : %p\n", sl2[8], &sl2[8])
	fmt.Printf("pointer of sl2[9] : %d , Address : %p\n", sl2[9], &sl2[9])
	fmt.Println("-------------------------Cut New name---------------------------")
	sl2Cut := sl2[5:]
	fmt.Println("sl2Cut[5:] : ", sl2Cut)
	fmt.Printf("pointer of sl2Cut[0] : %d, Address : %p\n", sl2Cut[0], &sl2Cut[0])
	fmt.Printf("pointer of sl2Cut[1] : %d, Address : %p\n", sl2Cut[1], &sl2Cut[1])
	fmt.Printf("pointer of sl2Cut[2] : %d, Address : %p\n", sl2Cut[2], &sl2Cut[2])
	fmt.Printf("pointer of sl2Cut[3] : %d, Address : %p\n", sl2Cut[3], &sl2Cut[3])
	fmt.Printf("pointer of sl2Cut[4] : %d, Address : %p\n", sl2Cut[4], &sl2Cut[4])
	fmt.Println("-------------------------Cut old name---------------------------")
	sl2 = sl2[5:]
	fmt.Println("sl2[5:] : ", sl2)
	fmt.Printf("pointer of sl2[0] : %d, Address : %p\n", sl2[0], &sl2[0])
	fmt.Printf("pointer of sl2[1] : %d, Address : %p\n", sl2[1], &sl2[1])
	fmt.Printf("pointer of sl2[2] : %d, Address : %p\n", sl2[2], &sl2[2])
	fmt.Printf("pointer of sl2[3] : %d, Address : %p\n", sl2[3], &sl2[3])
	fmt.Printf("pointer of sl2[4] : %d, Address : %p\n", sl2[4], &sl2[4])
	fmt.Println("----------------------------------------------------")

	sl3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("sl3 : ", sl3)
	sl3 = sl3[:]
	fmt.Println("sl3[:] : ", sl3)

}

func main() {
	fmt.Println("Hello Day3 Golang Slices")
	fmt.Println("----------------------------------------------------")
	sampleSlices()
	fmt.Println("----------------------------------------------------")
	sliceWithHightLow()
	fmt.Println("----------------------------------------------------")

}
