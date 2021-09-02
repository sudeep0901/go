package main

import (
	"fmt"
	"os"
)

type Employee struct {
	firstname string
	lastname  string
	city      string
	pincode   int
	permanent bool
	salary    float32
}

func (e *Employee) Value() float32 {

	value := float32((e.salary * 30) / 100)
	return value
}

func NewEmployee(fn string, ln string, city string, pc int, pe bool, salary float32) (*Employee, error) {

	if fn == "" {
		return nil, fmt.Errorf("firstname cannot be blank")
	}

	if ln == "" {
		return nil, fmt.Errorf("last name cannot be blank")
	}

	if city == "" {
		return nil, fmt.Errorf("city cannot be blank")
	}

	if pc == 0 {
		return nil, fmt.Errorf("pincode  cannot be blank")
	}

	if salary < 10000 {
		return nil, fmt.Errorf("salary cannot be less than 10000")
	}

	newemp := &Employee{
		firstname: fn,
		lastname:  ln,
		city:      city,
		pincode:   pc,
		permanent: pe,
		salary:    salary,
	}
	return newemp, nil
}
func main() {

	fmt.Print("Struct methods. . . . . . . .:")

	emp1 := Employee{"sudeep", "patel", "Ratlam", 457001, true, 10000.00}
	fmt.Println(emp1)

	fmt.Println(emp1.Value())

	emp2, err := NewEmployee("sudeep", "patel", "Ratlam", 457001, true, 10000.00)
	fmt.Print(emp2, err)

	emp2, err = NewEmployee("sudeep", "patel", "Ratlam", 457001, true, 0.00)
	fmt.Println(emp2, err)
	os.Exit(1)

}
