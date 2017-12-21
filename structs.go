package main

import "fmt"

type Details struct {
	name string
	age int
}

type empDB struct {
	name string
	age int
}


func main() {
	fmt.Println(Details{"nev", 21})
}

