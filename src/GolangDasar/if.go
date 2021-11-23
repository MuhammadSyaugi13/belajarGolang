package main

import "fmt"

func main() {


	// short if

	name := "ogi"

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
		}else{
		fmt.Println("Nama anda ", name)
	}

 }