package main

import "fmt"

type cellphone struct {
	name      string
	price     float32
	boomLevel int16
}

// -> value receiver: copy values, just read only
func (c cellphone) gulge() float32 {
	return c.price * float32(c.boomLevel)
}

func (c cellphone) amplify(index float32) float32 {
	return float32(c.boomLevel) * index
}

// -> pointer receiver: (ref pass) can modify struct properties
func (c *cellphone) changeName(newName string) {
	c.name = newName
}

func main() {
	iphone := cellphone{
		name:      "iphone",
		price:     6867.88,
		boomLevel: 1,
	}
	samsun := cellphone{
		name:      "samsun",
		price:     512.81,
		boomLevel: 100,
	}

	// p1xsel := cellphone{
	// 	name:      "p1xsel",
	// 	price:     4192.68,
	// 	boomLevel: 33,
	// }

	fmt.Println(iphone.price, samsun.boomLevel)
	fmt.Println(iphone.gulge(), samsun.gulge())
	fmt.Println(iphone.amplify(3.0) * iphone.price)

	// sabotage iphone
	iphone.changeName("ipone")

	fmt.Println(iphone)

}
