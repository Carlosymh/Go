package main

import "fmt"

func main() {

	var i int
	var num int

	fmt.Println("Ingresa Un Numero: ")
	fmt.Scanln(&num)

	fmt.Println("Tabla de Multiplicar del Numero", num)
	for i <= 10 {

		fmt.Println(num, "x", i, " = ", i*num)
		i++
	}
}
