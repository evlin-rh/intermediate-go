package print

import "fmt"

type User struct {
	Name string
	Age  int
}

func PrintMessage() {
	a := 1

	fmt.Println("This will join all arguments and print then. a =", a)
	fmt.Println("This will also insert space between args", "foo", "bar")
	fmt.Println("")

}

func PrintFMessage() {
	// use a '\\' to produce a single backlash
	fmt.Printf("A '\\' is called a backslash \n")

	fmt.Printf("The '%%' symbol is called a percent sign\n")
}

func PrintVerb() {
	s := "Hello world"

	// Print verb
	fmt.Printf("%s\n", s)

	// Print verb and wrap with quotes
	fmt.Printf("%q\n", s)

	// Print positive and negative numbers
	fmt.Printf("%d\n", 12345)
	fmt.Printf("%+d\n", -12345)

	// Padding
	fmt.Printf("'%5d'\n", 42)
	fmt.Printf("'%05d'\n", 42)
	fmt.Printf("'%5d'\n", 12345678)
	fmt.Printf("'%-5d'\n", 42)

	// Float
	fmt.Printf("'%f'\n", 1.2345)
	fmt.Printf("'%.2f'\n", 1.2345)

	//Structs
	u := User{
		Name: "Frank",
		Age:  23,
	}

	// %T -> Type
	fmt.Printf("%T\n", u)

	// %v -> Extended
	fmt.Printf("%v\n", u)

	// %#v -> Native Go-syntax format
	fmt.Printf("%#v\n", u)

	// Ordering
	fmt.Printf("%[4]s, %[3]s, %[2]s, %[1]s", "one", "two", "three", "four")

}
