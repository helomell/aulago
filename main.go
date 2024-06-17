package main 

import "fmt"

func main() {

	menu := map[string]float64{
		"pizza": 40.00,
		"suco": 12.50,
		"x-tudo": 22.90,
	}

	fmt.Println(menu["pizza"])
	for key, value:= range menu{ 
		fmt.Println(key, "R$", value)

	}

	contanova := novaConta("Astrubal")
	fmt.Println(contanova)

	contanova := novaConta("alexandro")
	fmt.Println(contanova2
}