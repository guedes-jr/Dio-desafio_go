package main

import "fmt"

func main() {
	var x [5]float64
	var total float64
	
	x[0] = 98.6
	x[1] = 99.5
	x[2] = 100.4
	x[3] = 101.3
	x[4] = 102.2
	
	fmt.Println(x)

	for i := 0; i < len(x); i++ {
		fmt.Printf("x[%d] = %.2f\n", i, x[i])
		total += x[i]
	}

	fmt.Printf("Total: %.2f\n", total)
	fmt.Printf("Média: %.2f\n", total/float64(len(x)))

}