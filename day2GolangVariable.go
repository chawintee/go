package main

import "fmt"

func initial() {
	fmt.Println("Hello Golang Day2")
}

func main() {
	initial()
	fmt.Println("Main Hello GOlang Day2")

	fmt.Println("----------------------------------------------------------------------")
	// variable golang Example
	// var name type
	// var <ชื่อตัวแปร> <ประเภทของตัวแปร>
	// var foo int

	var variableTypeBoolean bool = true
	var variableTypeInteger int = 65.000
	var variableTypefloat float64 = 54.56434567
	var variableTypeString string = "Hello day 2"
	var variableTypeFuction func() = func() {
		fmt.Println("In variableTypeFunction")
	}

	fmt.Println(variableTypeBoolean)
	fmt.Println(variableTypeInteger)
	fmt.Println(variableTypefloat)
	fmt.Println(variableTypeString)
	fmt.Println(variableTypeFuction)
	variableTypeFuction()

	// declare variable short hand example
	// var <ชื่อตัวแปร> = <ค่าของตัวแปร>
	// var bar = 99

	// <ชื่อตัวแปร> := <ค่าของตัวแปร>
	// zoo := 45

	fmt.Println("Variable short hand #1------------------------------------------------------------")

	var varBoolean = true
	var varInteger = 12
	var varFloat = 1.234
	var varString = "Hello Golang String"
	var varFuntion = func() {
		fmt.Println("Hello shorthand GOLANG function")
	}

	fmt.Println(varBoolean)
	fmt.Println(varInteger)
	fmt.Println(varFloat)
	fmt.Println(varString)
	fmt.Println(varFuntion)
	varFuntion()

	fmt.Println("Variable short hand #2----------------------------------------------------------")
	varrBoolean := true
	varrInteger := 88
	varrFloat := 3.5342
	varrSritng := "Hello GOLANG String #2"
	varrFunction := func() {
		fmt.Println("Hello shorthand variable declaration #2 function")
	}

	varrInteger = 23
	varrFloat = 43.5456

	const consttest = 54

	fmt.Println(varrBoolean)
	fmt.Println(varrInteger)
	fmt.Println(varrFloat)
	fmt.Println(varrSritng)
	fmt.Println(varrFunction)
	varrFunction()

	// consttest = 12
	fmt.Println(consttest)

}
