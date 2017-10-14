//Go by example.

package main

import(
  "fmt"
  "math/rand"
  "time"
)
func main(){
  fmt.Println("Hello world, the time is: ", time.Now())
  fmt.Println("My favorite number is", rand.Intn(10))
  }
