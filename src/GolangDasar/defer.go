package main

import "fmt"

func logging(){
	fmt.Println("Selesai memanggil logging")
}

func runApplication(value int){
	defer logging();
	fmt.Println("run App")
	result := 10 / value
	fmt.Println("result", result)
}

func main(){
	runApplication(0)
}