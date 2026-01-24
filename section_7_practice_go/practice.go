package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

var hobbies [3]string = [3]string{"Cooking", "Sports", "Reading"}

func main() {
	//1
	hobbies := [3]string{"Cooking", "Sports", "Reading"}
	fmt.Println(hobbies)

	//2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	//3
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	//4
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:]
	fmt.Println(mainHobbies)

	//5
	courseGoals := []string{"Learn go", "learn all the basics"}
	fmt.Println(courseGoals)

	//6
	courseGoals[1] = "Learn all the details!"
	courseGoals = append(courseGoals, "Learn advanced go")
	fmt.Println(courseGoals)

	//7
	products := []Product{
		Product{
			"first-product",
			"A first product",
			12.99},
		Product{
			"second",
			"Laptop",
			999.99},
	}

	fmt.Println(products)

	newProduct := Product{
		"third",
		"Monitor",
		199.99,
	}

	products = append(products, newProduct)
	fmt.Println(products)
}

func addHobby(newHobby string) [3]string {
	hobbies[0] = newHobby
	tmp := append(hobbies[1:3], "Gaming")[1:3]
	copy(hobbies[1:3], tmp)
	return hobbies
}
