package main

import "fmt"

func init1() {
	fmt.Println("int... Golang Day2 Loop...")
}

// ------------------------------------------------------------
// for loop
// for <กำหนดค่าเริ่มต้น>; <เงื่อนไข>; <หลังจากจบรอบ> {
//     ...
//     <คำสั่งการทำงาน>
//     ...
// }

func forLoop() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

// ---------------------------------------------------------------
// for loop + continue + break
// if else

func forLoopBreak() {
	for i := 0; i <= 100; i++ {
		if i == 5 {
			continue
		} else if i == 10 {
			break
		}
		fmt.Println(i)
	}
}

// -----------------------------------------------------
// while loop in GOLANG  "not clear"

//while in go
// for <เงื่อนไข> {
//     ...
//     <คำสั่งการทำงาน>
//     ...
// }

// for {
//    <คำสั่งการทำงาน>
// }

func whileInGoLang() {
	var i = 0
	for i <= 50 {
		fmt.Println(i)
		i++
	}
}

func whileInGoLangIfEleseContinueBreak() {
	i := 0
	for i < 100 {
		if i == 13 {
			continue
		}
		fmt.Println(i)
		i++
	}
	for {
		fmt.Println(i)
		break
	}
}

// --------------------------------------------------------------
// ในภาษาอื่น
// do {
//    <คำสั่งการทำงาน>
// } while(<เงื่อนไข>);
// ทางเลือกสำหรับคนอยากใช้
// for ok := true; ok; ok = <เงื่อนไข> {
//    <คำสั่งการทำงาน>
// }
// for {
//    <คำสั่งการทำงาน>
//    if !<เงื่อนไข> {
//      break
//    }
// }

func doWhileLoop_1() {
	i := 0
	for ok := true; ok; ok = i > 0 {
		fmt.Println(i)
		i--
	}
	fmt.Println("------------------------------------")
}

func doWhileLoop_2() {
	j := 0
	for {
		fmt.Println(j)
		if !(j > 0) {
			break
		}
	}
}

func main() {
	init1()
	fmt.Println("Hello GOLANG Loop...")
	// forLoop()
	// forLoopBreak()
	whileInGoLang()
	// whileInGoLangIfEleseContinueBreak()
	// doWhileLoop_1()
	// doWhileLoop_2()
}
