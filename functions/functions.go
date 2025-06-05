package functions

import (
	"fmt"
)

func PreferredFunction(greeting string, salutation string) {
	greeting = "hello"
	salutation = "goodbye"
	fmt.Printf(`It is better to have functions with all arguments \n 
				declare its type. This way is the preferred way`)
}

func NonNamedReturn1() (valid bool) {
	fmt.Println(valid)
	valid = true

	fmt.Println(valid)
	return
}

func NonNamedReturn() {
	NonNamedReturn1()
}
