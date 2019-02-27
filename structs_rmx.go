package main

import (
	"fmt";
	"encoding/json"
)

type Car struct {
		Make, Color 		string
		Model, Odo, Year 	int64
	}
	
var years map[int64]Car
//var makes map[string]years


func main() {
	fmt.Println("Hello, playground")
		
	var c Car
	c.Make = "Maruti"
	c.Model = 2000
	c.Odo = 100
	c.Year = 2019
	c.Color = "Red"
	
	//JSON Marshal and MarshalIndent
	data, _ := json.Marshal(c)
	human, _ := json.MarshalIndent(c, "", " ")
	
	fmt.Println("your new car: ",c)
	fmt.Printf("%s \n", data)
	fmt.Printf("%s \n", human)
	
	
	//Map
	years = make(map[int64]Car)
	years[1990]=Car{"Maruti", "Red", 2000, 100, 2019}
	fmt.Println(years[1990])
	
	
	
}
