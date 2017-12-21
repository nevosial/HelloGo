package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Details struct {
	name string
	age int
}

type empDB struct {
	name string
	age int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	
	
	//fmt.Println(Details{"nev", 21})
	
	//store records in the struct.
	v2 := empDB{"Max", 2331}	
	
	//Printing values.
	fmt.Println(v2)
	
	//Using dot
	fmt.Println(v2.name)
	fmt.Println(v2.age)
	
}








