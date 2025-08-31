package main

import "fmt"

func main() {
	// mapas
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
	// mapas no tienen orden
	delete(x, "key")
	fmt.Println(x["key"])
	// si no existe la llave, devuelve el valor cero del tipo
	v, ok := x["key"]
	fmt.Println(v, ok)

	// si no existe la llave, ok es false
	if !ok {
		fmt.Println("no existe")
	}

	// inicializar mapa con valores
	y := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(y)

	// agregar nuevo valor
	fmt.Println(y["foo"])
	fmt.Println(y["bar"])
	fmt.Println(y["baz"])
	fmt.Println(y)
}