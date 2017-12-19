//A defer statement defers the execution of a function until the surrounding function returns.
package main

import "fmt"

func main(){
  defer fmt.Println("Hello")

  fmt.Println("World")

  fmt.Println("Stacking defers")

  fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}
