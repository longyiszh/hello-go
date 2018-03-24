package main

import "fmt"
import "math/rand"
import "time"


func printRandom() {
  rand.Seed(time.Now().UTC().UnixNano())
  fmt.Println("Telling you a number [0 - 99]:", rand.Intn(100))
}

// func main() {
// 	printRandom()
// }