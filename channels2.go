//buffered channels.

package main

import "fmt"

func main() {
  // declaring the channel chq of type in and size 3.
	chq := make(chan int, 3)  
  
  //sending messages to the chn
	chq <- 27
	chq <- 21
	chq <- 88
  
  //printing the messages
  	v := <-chq
	fmt.Println(<-chq)
	//fmt.Println(<-chq)
  	fmt.Println(v)
	fmt.Println(<-chq)
}
