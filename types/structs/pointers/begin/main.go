// types/structs/pointers/begin/main.go
package main

import "fmt"

type author struct {
	first string
	last  string
}

// fullName returns the full name of the author
func (a author) fullName() string {
	return a.first + " " + a.last
}

// changeName changes the first and last name of the author
func (a *author) changeName(first, last string) {
	a.first = first
	a.last = last
}

func main() {
	a := author{
		first: "Mark",
		last:  "Twain",
	}

	fmt.Println(a.fullName())

	// call changeName to update name of author
	a.changeName("Charlie", "Chaplin")

	fmt.Println(a.fullName())
}
