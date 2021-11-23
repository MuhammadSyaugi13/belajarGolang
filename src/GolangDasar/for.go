package main

import "fmt"

func main() {

	//for biasa
	/*
	counter := 1
	
	for counter <= 10 {
		fmt.Println("perulangan ke - ", counter)
		counter++
	}
	
	*/
	
	// for dengan statement

	/*
	for count := 1; count <= 10; count++{
		fmt.Println("perulangan ke - ", count)
		
	}
	*/

	//  mengakses nilai slice dengan for

	/*
	slice := []string {"ogi", "bagir", "habil"}
	
	for i:=0; i < len(slice); i++{
		fmt.Println(slice[i])
	}
	*/
	
	// for dengan range
	/*
	slice := []string {"ogi", "bagir", "habil"}
	
	for i, value := range slice {
		fmt.Println("index ", i, "=", value)
	}
	*/

	//gunakan underscore jika ada variabel yang tidak dipakai
	/*
	for _, value := range slice {
		fmt.Println(value)
		} 
	*/

	// penggunaan for range pada map

	person := make(map [string]string)
	person["name"] = "ogi"
	person["title"] = "Backend Developer"

	// for range dengan map
	for key, value := range person {
		fmt.Println(key, "=", value)
	} 

	fmt.Println()
	//for range dengan statement underscore
	for _, value := range person {
		fmt.Println(value)
	} 



	
	



 }