package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	variadic("Header");
	variadic("Header", 1, 2);
	i, e := variadic("Header", 1, 2, 3, 4, 5);
	fmt.Println(i, e);

	if e != nil {
		os.Exit(1);
	}
}

// Go supports variadic parameters

func variadic(header string, vals ...int)  (int, error) {
	// out := make([]int, 0,len(vals))
	fmt.Println((vals));
	for _ ,v := range vals {
		fmt.Println(v * v);
	}

	return 1, errors.New("Multiple values returned by function");
}