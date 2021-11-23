package main

import "fmt"

func main() {

    // awal membuat map

    // person := map[string] string{
    //     "name" : "ogi",
    //     "address" : "bogor",
    // }

    // fmt.Println(person)
    // fmt.Println(person["name"])


    // deklarasi map
    var book map[string]string = make(map[string]string)
    book["title"] = "Belajar Golang"
    book["author"] = "Muhammad Saugi"
    book["ups"] = "salah"

    // tampilkan map
    fmt.Println(book)

    // panjang map
    fmt.Println(len(book))
    
    // hapus map
    delete(book, "ups")
    
    //tampilakan map setelah dihapus
    fmt.Println(book)


 }