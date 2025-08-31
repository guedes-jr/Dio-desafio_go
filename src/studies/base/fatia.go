package main

import "fmt"

func main() {
	arr := [7]float64{98.6, 99.5, 100.4, 101.3, 102.2, 103.1, 104.0}

	fatia := arr[1:4]
	fmt.Println(fatia)
	fatia2 := make([]float64, 2)
	copy(fatia2, fatia)
	fmt.Println(fatia2)
	fmt.Println(arr)
	fatia[0] = 0.0
	fmt.Println(fatia)
	fmt.Println(arr)
	fmt.Println(fatia2)
}