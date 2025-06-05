package maps

import (
	"fmt"
)

func SwitchStatement(temp int) {
	switch {
	case temp <= 32:
		fmt.Println("It is cold")
		fallthrough
	case temp >= 45 && temp < 90:
		fmt.Println("It is ok")
		fallthrough
	case temp >= 80:
		fmt.Println("It is hot")
	default:
		fmt.Println("Error")
	}
}
