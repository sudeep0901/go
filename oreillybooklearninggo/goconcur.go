package main

import (
	"fmt"
)

func main() {
	fmt.Println("Program started............................");
	ch := make(chan string);
	fmt.Println(ch);
	// writing to channel
	ch <- "Sudeep Patel";
	// Reading channel values
	a := <-ch;
	fmt.Println(a)
}

func count(thing string) {
	
}