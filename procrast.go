package main

import "fmt"
import "time"
import "sync"

var wGroup sync.WaitGroup

func procrast(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str)
		time.Sleep(time.Millisecond * 100)
	}
}

func procrastD(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str)
		time.Sleep(time.Millisecond * 100)
	}

	defer wGroup.Done() // defer: begin to run when other surrounding statements complete or panic out
	// Defer behaves like a stack. Statements in multiple "defer"s follow LIFO order of exec
}

func procrastN(str string) {

	defer safeWGDone()

	for i := 0; i < 3; i++ {
		fmt.Println(str)
		time.Sleep(time.Millisecond * 100)
		if i == 2 {
			makeTrouble()
		}
	}

}

func makeTrouble() {
	panic("Trouble Starts!")
}

func safeWGDone() {

	defer wGroup.Done() // defer: begin to run when other surrounding statements complete or panic out
	// Defer behaves like a stack. Statements in multiple "defer"s follow LIFO order of exec

	recoverer := recover()

	if recoverer != nil {
		fmt.Println("Exception handled: ", recoverer)
	}

}

func freeStyle() {
	go procrast("beep") // async

	fmt.Println("Don")

	go procrast("boop") // async

	time.Sleep(time.Second)
}

func withRoutines() {

	wGroup.Add(1)
	go procrastD("Baa") // async

	fmt.Println("Boom")

	wGroup.Add(1)
	go procrastD("Werr") // async

	wGroup.Wait()

}

func withRoutinesSafe() {

	wGroup.Add(1)
	go procrastN("Von") // async

	fmt.Println("Beep")

	wGroup.Add(1)
	go procrastN("Dann") // async

	wGroup.Wait()

}

func mainpr() {
	withRoutinesSafe()
}
