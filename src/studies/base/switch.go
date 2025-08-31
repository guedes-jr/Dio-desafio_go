package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		switch i {
			case 0: fmt.Printf("zero\n")
			case 1: fmt.Printf("um\n") 
			case 2: fmt.Printf("dois\n") 
			case 3: fmt.Printf("três\n") 
			case 4: fmt.Printf("quatro\n") 
			case 5: fmt.Printf("cinco\n") 
			case 6: fmt.Printf("seis\n") 
			case 7: fmt.Printf("sete\n") 
			case 8: fmt.Printf("oito\n") 
			case 9: fmt.Printf("nove\n") 
			case 10: fmt.Printf("dez\n") 
		}
	}
}