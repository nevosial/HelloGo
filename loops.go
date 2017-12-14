package main

import "fmt"

func main(){

  fmt.Println("Type 1")

  i := 1
  for i <=3 {
    fmt.Println(i)
    i = i + 1
  }

  fmt.Println("Type 2")
  for j := 1; j <=3; j++ {
    fmt.Println(j)
  }

  fmt.Println("Type 3")
  for{
    fmt.Println("loop")
    break
  }

  fmt.Println("Odd Numbers between 0 and 8")
  for n := 0; n <= 8; n++ {
    if n%2 == 0 {
      continue
    }
    fmt.Println(n)
  }
  
}
