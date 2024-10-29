package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	user := User{"pallav", "pallav@go.dev", true, 16}
	fmt.Println(user)
	fmt.Printf("pallav details are: %+v\n", user)
	fmt.Printf("Name is %v and email is %v.", user.Name, user.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
