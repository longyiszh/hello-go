package main

import "fmt"

func main1() {
	fmt.Println("hello-go")

	// pointer type
	a := 15
	pt := &a

	fmt.Println(a)

	fmt.Println(pt, *pt)

}
