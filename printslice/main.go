package main

import "fmt"

func main() {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, l := range list {
		if l%2 != 0 {
			fmt.Printf("%v is Odd\n", i)
		} else {
			fmt.Printf("%v is Even\n", i)
		}
	}
}
