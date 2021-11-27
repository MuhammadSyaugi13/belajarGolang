package main

import "fmt"

type Costumer struct {
	name, address string
	age int
}

//struct method
func (costumer Costumer) sayHello(name string){
	fmt.Println("Hello ", name, " my name is ", costumer.name)
}



func main(){
	var ogi Costumer

	ogi.name = "Muhammad Syaugi"
	ogi.address = "Bogor"
	ogi.age = 22
	ogi.sayHello("adi")

	//fmt.Println(ogi)
	
	//struct literals
	// pertama
	joni := Costumer{
		name: "joni Iskandar",
		age: 22,
		address: "bogor",
	}
	
	fmt.Println(joni)
	
	// kedua
	//isi didalam kurung kurawa urutannya tidak boleh beda
	budi := Costumer{"ogi", "Jakarta", 22}

	fmt.Println(budi)


}























