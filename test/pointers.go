package main

import "fmt"

func main() {
	var a string = "hello world"
	fmt.Println(a) // output: hello world.

	//"&" says: "get the address of the variable that follows".
	fmt.Println(&a) // output: 0xc000010070.

	//"*" says: you are declaring that this variable HOLDS A MEMORY
	//ADDRESS TO A STRING, or int, or whatever type follows "*".
	//For example, "var a *int" declares that the variable "a" holds a
	//memory address(pointer) to an int datatype.
	var b *string = &a //declare "b" as type "pointer to string".

	fmt.Println(b)  // output: 0xc000010070
	fmt.Println(&b) // output: 0xc0000b4020

	//"*"" says: "give me whatever the variable that follows is pointing to."
	//So your variable better be an actual memory address, otherwise you’ll get an error.
	//It tells the system to use the value as a pointer and return whatever is at that address.
	fmt.Println(*b) // output: hello world
}
