package main

import "fmt"

func main() {

	// break
	fmt.Println("Break")
	for i := 0; i < 10; i++{

		if i == 5 {
			break
		}

		fmt. Println("i = ", i)
	}
	fmt. Println()

	// continue 
	fmt.Println("Continue")
	for i := 0; i < 10; i++{

		if i%2 == 0 {
			continue
		}

		fmt. Println("i = ", i)
	}
}