package main

import (
	"fmt"
)

func main() {

	fmt.Println("Main function starting.......");
	var x int32 = 10;
	var name string = "sudeep";
	pointerX := &x;

	fmt.Println(pointerX);

	var pointerZ  *string
	fmt.Println(pointerZ);

	pointerZ = &name;
	fmt.Println(pointerZ);

	fmt.Println(*pointerX);
	var pointerZPanic  *string
	fmt.Println(pointerZPanic==nil)
	// fmt.Println(*pointerZPanic)

	var n = new(int)
	fmt.Println(n, *n);

	type person struct {
		fname string
		lname *string
		city string
	}
	p := person {
		fname: "Sudeep",
		lname: stringp("Patel"),
		city: "ratlam",
	}

	ptr := stringp(name)
	fmt.Println(ptr, p)
	fmt.Println(p.lname, *p.lname)
	
	var f *int;
	fmt.Println(f);
	failedupdate(f);
	fmt.Println(f);
	
	type MailCategory int

	const (
		Uncategorized MailCategory = iota
		Personal
		Spam
		Social
		Advertisements
	)
}	
func stringp(s string) *string {
	return &s
}

func failedupdate(f *int) {
	x := 10;
	f = &x;

	fmt.Println(f)
}

