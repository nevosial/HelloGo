package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func pickbox() {
	fmt.Println("Move box 1.")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Move box 2.")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Move box 3.")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Move box 4.")
}

func main() {
	go pickbox()
	go say("world")
	say("hello")

}

/*
Move box 1.
hello
world
Move box 2.
world
hello
Move box 3.
Move box 4.
hello
world
world
hello
world
hello
*/
