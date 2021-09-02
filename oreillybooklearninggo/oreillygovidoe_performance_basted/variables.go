package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// var is zero value
	var a int
	var b string
	var c float64
	var d bool
	fmt.Println("Hello, playground")

	fmt.Println("var a int  %X [%v] \n", a, a)
	fmt.Println("var b int \t %T [%v] \n", b, b)
	fmt.Println("var c int \t %T [%v] \n", c, c)
	fmt.Println("var d int \t %T [%v] \n", d, d)
	h := fmt.Sprintf("%X",b)
	fmt.Print(h)


	type example struct {
		flag bool
		counter int64
		pi float32
	}

	type example2 struct {
		counter int64
		pi float32

		flag bool
	}
	// padding and alignments - padding after flag to keep alignment boundary in tact

	// solution order field largest to smallents
	var e1 example
	fmt.Println(e1)
	var e1size = unsafe.Sizeof(e1)
	fmt.Println(e1size)
	var e2 example2

	var e2size = unsafe.Sizeof(e2)
	fmt.Println(e2size)
	fmt.Printf("e2: %v\n", e2)
	fmt.Printf("e2: %v\n", e2.counter)

	// annonymous struct
	 anno := struct {
		flag bool
		name string
		age int
	}{
		flag: true,
		name: "Sudeep",
		age: 22, 
	}

	fmt.Println(anno)


	type bill struct {
		counter int64
		pi float32

		flag bool
	}

	type alice struct {
		counter int64
		pi float32

		flag bool
	}

	var bl bill
	var al alice
	
	fmt.Println(bl)
	fmt.Println(al)
	// no implicit conversion allowed in golang
	// bl = al
	// to do abovedo conversion

	// Do explicit conversion like below
	bl = bill(al)

	// when type is literal implicit conversioncons

}
