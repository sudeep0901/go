package main

import (
	"fmt"
)

func main() {

	fmt.Println("Main function starting compostion.......");
	e := Employee {
		Name: "Sudeep",
		ID: "1234",
	}

	fmt.Println(e.Name, e.ID);
	m := Manager{
		Employee: e,
		Reports: []Employee{},
	}
	fmt.Println(m)
}


type Employee struct {
	Name string
	ID string
}
type Manager struct {
	Employee
	Reports []Employee
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID);
}

// func (m Manager) FindNewEmployee() []Employee {
// 	// reutrn new([]);
// }	