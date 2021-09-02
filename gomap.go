package main

import "fmt"

func main() {
	fmt.Print("Starting the program")

	stocks := map[string]string{
		"INDIA":       "DELHI",
		"USA":         "WASHINTON D.C.",
		"SOUTH KOREA": "SEOUL",
	}

	fmt.Println(stocks)
	fmt.Println(len(stocks))

	fmt.Println(stocks["INDIA"])

	value, ok := stocks["INDIA"]
	fmt.Println(value, ok)
	if !ok {
		fmt.Println("Data found")
	} else {
		fmt.Println("Data not found")

	}
	value, ok = stocks["INDIAA"]
	fmt.Println(value, ok)

	for key, value := range stocks {
		fmt.Println(key, value)
	}
}
