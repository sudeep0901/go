package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var x int = 10;
	var y = [3]int{1, 2, 3}
	// var xy = [4]int{1, 2, 3,4}
	// xy = y;
	var arr1 = [...]int{4, 5, 6}

	var arr2 = [12]int{0,1:5,2:10}
	fmt.Println(x);
	fmt.Println(y)
	fmt.Println(arr1);
	fmt.Println(len(arr1));
	fmt.Println(arr2);
	// fmt.Println(y == xy);
	fmt.Println(y[2]);

	// Using […] makes an array. Using [] makes a slice.
	// Go slices

	var slice = []int{1, 2, 3,4, 5, 10, 100: 1}
	var slice1 = []int{1, 2, 3}

	fmt.Println(slice);
	slice = append(slice, 100);
	fmt.Println(append(slice, 1001));
	fmt.Println(slice);
	fmt.Println(len(slice));
	var slc []int
	fmt.Println(slc);
	fmt.Println(slc == nil);
	// fmt.Println(slice == slice1);
	fmt.Println(len(slc));
	fmt.Println(len(slice1));
	fmt.Println(cap(slice));
	xslice := make([]int, 5, 10)
	fmt.Println(cap(xslice), xslice);

	var name = "sudeep patel";
	fmt.Println(name[1:3]);

	var a rune    = 'x'
	var s string  = string(a)
	var b byte    = 'y'
	var s2 string = string(b)

	fmt.Println(a, s, b, s2);

	var nilMap map[string]int;

	fmt.Println(nilMap,  len(nilMap));
	totalWins := map[string]int{}
	fmt.Println(totalWins,  len(totalWins));

	teams := map[string][]string {
		"Orcas": []string{"Fred", "Ralph", "Bijou"},
		"Lions": []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}

	fmt.Println(teams);

	ages := make(map[int][]string, 10);
	fmt.Println(len(ages));
	v, ok := teams["Orcas"];
	fmt.Println(v, ok);

	v, ok = teams["orcas"];
	fmt.Println(v, ok);

	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	for _, v:= range vals {
		fmt.Println(v);
	}

}

