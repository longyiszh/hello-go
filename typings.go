package main

import "fmt"

func multipleStrings(str1,str2 string) (string,string) {
	return str1, str2
}

func add(x,y float64) float64 {
	return x + y
}

func main() {
	var str1 = "Hello"

	str2 := "  " + "go go"

	fmt.Println(multipleStrings(str1, str2))

	var num1 int = 32
	var num2 float32 = 5.8877

	x := float64(num1)
	y := float64(num2)

	fmt.Println(x, y)
	fmt.Println(add(x, y))

}