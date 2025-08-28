// Seu programa deve imprimir numeros pares de 0 a 100, porem quando o numero for divisivel 
// por 3 deve imprimir "Pin", caso seja divisivel por 5 deve imprimir "Pan".	
package main

import "fmt"

func main() {
	for i := 0; i <= 100; i ++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PinPan")
		} else if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}
