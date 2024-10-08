package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	pallav := User{"pallav", "pallav@go.dev", true, 16}
	fmt.Println(pallav)
	fmt.Printf("pallav details are: %+v\n", pallav)
	fmt.Printf("Name is %v and email is %v.", pallav.Name, pallav.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
