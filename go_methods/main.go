package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	pallav := User{"pallav", "pallav@go.dev", true, 16}
	fmt.Println(pallav)
	fmt.Printf("pallav details are: %+v\n", pallav)
	fmt.Printf("Name is %v and email is %v.\n", pallav.Name, pallav.Email)
	pallav.GetStatus()
	pallav.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", pallav.Name, pallav.Email)

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
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
