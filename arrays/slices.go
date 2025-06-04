package arrays

import "fmt"

func LenCap() {
	var s1 []int

	hat := cap(s1)
	for i := 0; i <= 1000; i++ {
		s1 = append(s1, i)
		c := cap(s1)
		if c != hat {
			fmt.Println(hat, c)
		}
		hat = c
	}
}

func SliceMethods() {
	a := []string{}        // Declared to an empty string slice
	var b []string         // Declared as a slice of string
	c := make([]string, 0) // Created using a make function

	fmt.Println(a, b, c)
}

func IncorrectAppend() {
	a := make([]string, 2)
	a = append(a, "foo", "bar")

	fmt.Printf("%q\n", a)
}

func Copy() {
	names := []string{"Hello", "Goodbye", "Salut"}
	subset := make([]string, 2)
	copy(subset, names)

	fmt.Println(subset)
}

func ArrToSlice() {
	names := [4]string{"A", "B", "C", "D"}

	slicesOnly(names[:])
}

func slicesOnly(name []string) {
	for _, v := range name {
		fmt.Println(v)
	}
}

func BreakContinue() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		if i == 2 {
			continue
		}

		fmt.Println(i)
	}
}
