package main

import (
	"fmt"
)

func main() {

	fmt.Println("Main function starting.......");
		type MailCategory int

	const (
		Uncategorized MailCategory = 100 << iota
		Personal
		Spam
		Social
		Advertisements
	)

	fmt.Println(Personal);
	fmt.Println(Spam);


}	