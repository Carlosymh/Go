package main

import "fmt"

func main() {
	var Comida string
	fmt.Println("¿Que hay de Comer? ")
	fmt.Scanln(&Comida)
	if Comida == "Pizza" || Comida == "pizza" || Comida == "Pizzas" || Comida == "pizzas" {
		fmt.Println("Siii  Pizza!! :D")
	} else if Comida == "Hamburguesa" || Comida == "hamburguesa" || Comida == "Hamburguesas" || Comida == "hamburguesas" {
		fmt.Println("Siii  Hamburgesas!! :D")
	} else if Comida == "Tamales" || Comida == "tamales" {
		fmt.Println("Bueno Lo Aceptare :| ")
	} else {
		fmt.Println("Yo Queria Pizza :C")
	}

}
