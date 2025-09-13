package main

import "fmt"

//go garbage collector in this case when use the function the variable age is copied to the function stack frame
// and when the function ends the stack frame is deleted and the variable age is lost
// if we want to modify the variable age in the function we need to use pointers

func main() {
	age := 32 /// regular variable

	var agePointer *int

	agePointer = &age // pointer to the variable age
	// fmt.Println("Age Pointer Value: ", *agePointer) // dereference the pointer to get the value

	fmt.Println("Age: ", *agePointer)
	editAgeToAdultYears(agePointer)
	//El valor inicial de age con el anterior metodo getAdultYears con el valor de 32
	//A verificar el metodo te hace la resta de 18 y te da 14
	//La diferencia grande es que todo esta en la variable inicial de la memoria de la variable age
	//Esto es gracias a los punteros
	//Si no se usaran punteros el valor de age seguiria siendo 32
	//Porque la funcion getAdultYears recibe una copia del valor de age
	//Y no el valor original
	//Por lo tanto cualquier cambio que se haga en la funcion no afectara al valor original
	//Ya que solo se esta trabajando con una copia del valor original
	//Y cuando la funcion termina la copia se pierde y el valor original sigue siendo el mismo
	//Ahora entendido esto le ponemos el nombre de la variable editAgeToAdultYears
	fmt.Println(age)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
