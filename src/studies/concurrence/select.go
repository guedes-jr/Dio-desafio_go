package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second *1)
		c1 <- "um"
	}
	go func() {
		time.Sleep(time.Second *2)
		c2 <- "dois"
	}

	for i := 0; i < 10; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("receba", msg1)
		case msg2 := <-c2:
			fmt.Println("receba", msg2)
		}
	}
}
