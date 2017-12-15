package main

import "fmt"

func add(a int, b int) int {
  return a + b
}

func main(){

  res := add(4, 3)
  fmt.Println("4+3= ",res)
}
