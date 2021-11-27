package main

import "fmt"

func endApp(){
	
	message := recover()
	if message != nil {
		fmt.Println("pesan error : ", message)
		
	}
	fmt.Println("pesan selesai")
}

func runApp(error bool){
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}

	fmt.Println("aplikasi berjalan")
}

func main(){
	runApp(true)
	fmt.Println("ogi")
}