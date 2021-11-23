package main

import "fmt"

//function with return multiple value
func getFullName(firstName string, lastName string) (string, string){
	
	return firstName, lastName
	
}

//function with Named return multiple value
func getNameReturn() (firstName string, middleName string, lastName string)  {
	firstName = "Muhammad"
	middleName = "Syaugi"
	lastName = "Quthban"
	return
}

// function variadic
func sumAll(numbers ...int) int{
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

//function value & function as parameter

type Filter func(string) string
func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)

}

func spamFilter(name string) string  {
	if name == "Anjing" {
		return "..."
	}
	return name
}

//anonymous function
type Blacklist func(string) bool 
func registerUser(name string, blacklist Blacklist)  {
	if blacklist(name) {
		fmt.Println("you are Blocked ", name)
	}else{
		fmt.Println("Welcome ", name)
	}
	
}

//recursive function
/*intinya func yg memanggil dirinya sendiri*/

func recursiveFactorial (value int) int {
	if value == 1 {
		return 1
	}else {
		return value * recursiveFactorial(value - 1)
	}
} 

func main(){


	//recursive function
	fmt.Println(recursiveFactorial(5))
	

	//multple value
	/*
	println("return multiple value")
	
	namaAwal, namaAkhir := getFullName("Muhammad", "Syaugi")
	fmt.Println("Nama Awal : ", namaAwal, "\nNama Akhir : ", namaAkhir)
	
	
	fmt.Println()
	
	//named return value
	println("return Name Return Value")
	
	a, b, c := getNameReturn()
	fmt.Println("Nama Awal : ", a, "\n",
	"Nama tengah : ", b, "\n",
	"Nama akhir : ", c,
	)
	*/
				
				
	// output function variadic
	/*
	println("return function variadic")
	
	total := sumAll(10, 4, 13)
	fmt.Println(total)
	
	println()
	println("return function variadic dengan parameter slice")
	
	slice := []int{10,20,30,40,50}
	total = sumAll(slice...)
	fmt.Println(total)
	*/

	// functiion value & function as parameter
    /*
	sayHelloWithFilter("ogi", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
	*/

	//function anonymous

	/*
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})


	//funtion value
	blacklist := func (name string) bool  {
		//cek apakah nilai parameter sama seperti Admin, jika kembalikan nilai true 
		return name == "admin"
	}
	
	registerUser("admin", blacklist)
	*/	


	
	

}