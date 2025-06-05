package maps

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func Maps() {
	var users map[string]string
	fmt.Println(users)

	users = make(map[string]string)
	users["1"] = "Kurt"
	users["2"] = "Brian"
	fmt.Println(users)
}

func DeleteMap() {
	users := map[string]string{}

	users["1"] = "evan"
	users["2"] = "evan"
	users["3"] = "evan"

	delete(users, "2")
	delete(users, "What")

	fmt.Println(users)
}

func BoolKey() {
	users := map[string]string{}

	users["1"] = "evan"
	users["2"] = "evan"
	users["3"] = "evan"

	value, test := users["hello"]
	if !test {
		fmt.Printf("Key %s not found\n", value)
		os.Exit(1)
	}

	fmt.Println(users)
}

func CountMap() {
	counts := map[string]int{}
	counts2 := map[string]int{}

	text := "The string is called the string"
	newText := strings.Fields(strings.ToLower(text))

	for _, v := range newText {
		counts[v] = counts[v] + 1
	}

	for _, v := range newText {
		counts2[v]++
	}

	fmt.Println(counts)
	fmt.Println(counts2)
}

type User struct {
	ID   int
	Name string
}

func Struct() {
	data := map[int]User{}

	user := User{ID: 1, Name: "Kurt"}

	data[1] = user
	user2 := data[1]
	user2.Name = "Ann"
	data[1] = user2

	fmt.Println(data)
}

func SortKeys() {
	months := make(map[int]string)
	months[1] = "Jan"
	months[2] = "Feb"
	months[3] = "Mar"
	months[4] = "Apr"
	months[5] = "May"
	months[6] = "Jun"
	months[7] = "Jul"

	keys := make([]int, 0)

	for k := range months {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for _, v := range keys {
		fmt.Printf("%q\n", months[v])
	}

}
