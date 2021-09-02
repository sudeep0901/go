package main

import (
	"fmt"

	"sudeep0901/gomodules/math"
	"sudeep0901/gomodules/formatter"
	
)

func main() {
	num := math.Double(2)
    output := print.Format(num)
    fmt.Println(output)
}

