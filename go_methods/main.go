package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	user := User{"pallav", "pallavps17@gmail.com", true, 16}
	fmt.Println(user)
	fmt.Printf("pallav details are: %+v\n", user)
	fmt.Printf("Name is %v and email is %v.\n", user.Name, user.Email)
	user.GetStatus()
	user.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", user.Name, user.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@gmail.com"
	fmt.Println("Email of this user is: ", u.Email)
}
