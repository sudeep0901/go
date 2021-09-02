package main

import (
	"errors"
	"fmt"
	"log"
)


func main() {
	fmt.Println("starting go lang errors");
	log.SetFlags(6);
	log.SetPrefix("greetings:\t");
	message, err:= Hello("");

	if err != nil {
		log.Fatal(err);
	}

	// If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}


func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Empty Name");
		
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}
