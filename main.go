package main

import (
	"fmt"

	arrays "github.com/abozadev/go-playground/arrays"
)

func main() {
	res := arrays.IsUnique("Helo")
	if res {
		fmt.Println("Is Unique")
	} else {
		fmt.Println("Not Unique")
	}
}
