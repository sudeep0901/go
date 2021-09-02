package main

import (
	"fmt"
)

type user struct {
	name string
	email string
}
func main() {
	fmt.Println("factory functions")
	u := CreateUser1()	
	fmt.Println("V1",u, &u)
	var u2 *user
	u2 = CreateUser2()	
	fmt.Println("V2",u2, &u2)


	u3 := CreateUser3()
	fmt.Println(u3, &u3)
}
// factory function
func CreateUser1() user {
	u:= user {
		name: "sudeep",
		email: "sudeep.tech.patel@gmail.com",
	}
	fmt.Println("V1",u, &u)

	return u
}

// factory function
// Escape Analysis, moves it to heap escapes from stack, compiler makes static code analysis
func CreateUser2() *user {
	// literal constuction
	u:= user {
		name: "sudeep1",
		email: "sudeep1.tech.patel@gmail.com",
	}
	fmt.Println(u, &u)
	// u variable will not be created in stack but in heap which is shared memory
	return &u
}


func CreateUser3() *user {
	// literal constuction
	u:= &user {
		name: "sudeep1",
		email: "sudeep1.tech.patel@gmail.com",
	}
	fmt.Println(u, &u)
	// u variable will not be created in stack but in heap which is shared memory
	return u
}