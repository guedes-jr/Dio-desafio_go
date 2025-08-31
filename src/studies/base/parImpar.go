package main

import "fmt"

func main() {
	for i:= 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d é par\n", i)
		} else {
			fmt.Printf("%d é ímpar\n", i)
		}
	}
}