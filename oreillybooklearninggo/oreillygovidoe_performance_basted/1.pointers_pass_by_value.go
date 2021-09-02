package main

import (
	"fmt"
)

func main() {
	fmt.Println("Frame memory allocated from go and varialbes defined into it")	
	c := 10

	fmt.Println(c, &c)
	increment(c)
	fmt.Println(c, &c)

}

func increment(v int) {
	v = v + 1
	v++
	fmt.Println("new slice of stack frame will be given and isolation")
	fmt.Println(v)
}