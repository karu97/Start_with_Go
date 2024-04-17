package main

import "fmt"

func main() {

	//strings
	var Go_Str string = "I want to be the best software engineer in the world"
	var Go_str_01 = "Praveen Karunanayake"

	fmt.Println(Go_Str)
	fmt.Println(Go_str_01)

	Go_str_02 := "+94771703811"
	fmt.Println(Go_str_02)

	//int

	var Go_Int int = 123456
	var Go_Int_01 = 678910
	var Go_Int_02 = 111213141516

	fmt.Println(Go_Int, Go_Int_01, Go_Int_02)

	//bits & memory

	var go_memory int8 = 25
	var go_memory_01 int16 = 278

	fmt.Println(go_memory, go_memory_01)

	var go_memory_02 float64 = 1.00012121

	fmt.Println(go_memory_02)

	//print

	fmt.Print("55/2,  \n")
	fmt.Print("Church Road, \n")
	fmt.Print("Gampaha.")

	age := 25
	name := "Praveen"

	fmt.Println("My name is :", name)
	fmt.Println("My age is :", age)

	//printf (Format Strings)

	fmt.Printf("My name is %v and My age is %v", name, age)

}
