// concurrency/channel-non-blocking/begin/main.go
package main

import (
	"fmt"
	"time"
)
func main() {
	// declare a signal channel to read from
	readChan := make(chan struct{})

	// use a default case in select to prevent blocking
  select {
	  case <- readChan:
	    fmt.Println("Reading from readChan")
    default:
      fmt.Println("No signal found")
  }

	timeChan1 := time.After(200 * time.Millisecond)
	timeChan2 := time.After(400 * time.Millisecond)

	for {
		select {
  		case <- timeChan1:
				fmt.Println("Received from timeChan1")
				
			case <- timeChan2:
				fmt.Println("Received from timeChan2")
				return
			default:
				fmt.Println("Received no signal")
				time.Sleep(50 * time.Millisecond)
		}
	}
}
