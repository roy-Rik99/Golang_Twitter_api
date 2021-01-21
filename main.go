package main

import (
	"Projects/twitterapi"
	"fmt"
)

func main() {

	var first, last string
	fmt.Print("Enter first?")
	fmt.Scanf("%s", &first)
	fmt.Print("Enter last?")
	fmt.Scanf("%s", &last)
	name := (first + " " + last)
	fmt.Println(name)
	twitterapi.RequestUserDetails(name)
	fmt.Printf("\nUser Details :---\n\t%#v\n\n", twitterapi.UserInfo)
}
