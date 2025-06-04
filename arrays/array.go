package arrays

import "fmt"

func AppendSlices() {
	var names []string
	names = append(names, "Kurt")

	var names2 []string
	names2 = append(names2, "Brian")

	/* This does not work
	names = append(names, names2)
	*/

	/* Loops was previously used
	for _, name := range names2 {
		names = append(names, name)
	}
	*/

	// New Format
	names = append(names, names2...)
	fmt.Println(names)
}
