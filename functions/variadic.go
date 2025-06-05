package functions

import (
	"fmt"
)

func MultipleVariadic(num int, list ...string) {
	for _, n := range list {
		fmt.Printf("%d. %s\n", num, n)
		num++
	}
}

func VariadicSlice(list ...string) {
	for _, n := range list {
		fmt.Printf("%s\n", n)
	}
}
