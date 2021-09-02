package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go structs");

	type person struct {
		name string
		age int
		pet string
	}

	var sudeep person;
	fmt.Println(sudeep);
	fmt.Println(sudeep.pet == "");

	sudeep.name ="Sudeep"
	sudeep.age = 26;
	fmt.Println(sudeep);
	fmt.Println(sudeep.pet == "");

	bob := person{};
	fmt.Println(bob);
	beth := person{
		age:  30,
		name: "Beth",
	}

	fmt.Println(beth);
	
	// Anonymous Structs

	var person1 struct {
		name string
		age  int
		pet  string
	}
	
	person1.name = "bob"
	person1.age = 50
	person1.pet = "dog"
	

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}

	fmt.Println(pet);
	fmt.Println(person1);

}

