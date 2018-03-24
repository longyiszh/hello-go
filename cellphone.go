package main

import "fmt"

type cellphone struct {
	name string
	price float32
	boomLevel int16
}

func main() {
	iphone := cellphone {
		name: "iphone",
		price: 6867.88,
		boomLevel: 1,
	}
	samsun := cellphone {
		name: "samsun",
		price: 19512.81,
		boomLevel: 100,
	}

	fmt.Println(iphone.price, samsun.boomLevel)
}