package main

import "fmt"

func main(){

  //creating new map using builtin make
  m := make(map[string]int)

  m["a1"]= 27
  m["a2"]= 89

  fmt.Println("map: ", m)

  //getting the value for a key a1.
  v1 := m["a1"]
  fmt.Println("value of first key: ",v1)
  //length of the map
  fmt.Println("length of map: ", len(m))

  //deleting values
  delete(m, "a2")
  fmt.Println("New Map after deleting key/value a2: ",m)


  //finding if given key is present in the map.
  _, prs := m["a2"]
  fmt.Println("check if a2 is present:", prs)

  //initializing map in the same line.
  n := map[string]int{"foo":1, "bar":2}
  fmt.Println("initializing type 2: ",n)

}
