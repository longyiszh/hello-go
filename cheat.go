package main

import "fmt"
import "sync"
import "math/rand"
import "time"

var wGroup sync.WaitGroup

func random() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(100)
}

func cheat(c chan int, val int) {
	defer wGroup.Done()
	c <- val*val + 1
}

func main() {

	scoreChannel := make(chan int, 10) // 10: channel buffer size

	for i := 0; i < 10; i++ {
		wGroup.Add(1)
		go cheat(scoreChannel, random())
		time.Sleep(time.Millisecond * 100)
	}

	wGroup.Wait()

	// Before reading from buffer channel we need to close it (since we used async function)
	close(scoreChannel)

	for score := range scoreChannel {
		fmt.Printf("You get %d / 100 in the final exam!\n", score)
	}

}
