package main

import "fmt"

type cellphone struct {
  name string
  price float32
  boomLevel int16
}

func (c cellphone) gulge() float32 {
  return c.price * float32(c.boomLevel)
}

func (c cellphone) amplify(index float32) float32 {
  return float32(c.boomLevel) * index
}

func main() {
  iphone := cellphone {
    name: "iphone",
    price: 6867.88,
    boomLevel: 1,
  }
  samsun := cellphone {
    name: "samsun",
    price: 512.81,
    boomLevel: 100,
  }

  fmt.Println(iphone.price, samsun.boomLevel)
  fmt.Println(iphone.gulge(), samsun.gulge())
  fmt.Println(iphone.amplify(3.0) * iphone.price)
}