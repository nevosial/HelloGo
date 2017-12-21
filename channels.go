//simple channel

package main

import "fmt"

func main() {
  // declaring the channel chq of type in and size 2.
	chq := make(chan int, 2)  
  
  //sending messages to the chn
	chq <- 1
	chq <- 2
  
 
  v := <-chq
    
  //printing the messages(Note the order)
	fmt.Println(<-chq)
  fmt.Println(v)
  
}
