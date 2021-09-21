package main

import (
	"fmt"
)


type user struct {
	name string
	email string
}
func main() {
	fmt.Println("Starting main funcion")
	sudeep := user{"Sudeep", "sudeep.agile@gmail.com"}
	fmt.Println(sudeep)

	sudeep.changeEmail("sudeep.tech.patel@gmail.com")

	fmt.Println(sudeep)
}


// receiver u is 
// this is method
func (u *user) notify() {
	fmt.Printf("sending email to user %s<%s> \n", u.name, u.email)
}

func (u *user) changeEmail(eml string) {
	fmt.Printf("sending email to user %s<%s> \n", u.name, u.email)
	u.email = eml
}