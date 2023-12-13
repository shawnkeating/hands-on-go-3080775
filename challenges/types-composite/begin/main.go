// challenges/types-composite/begin/main.go
package main

import (
	"fmt"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title  string
	author author
}

// define a library type to track a list of books library struct {
type library map[string][]book

// define addBook to add a book to the library
func (l library) addBook(b book) {
	l[b.author.name] = append(l[b.author.name], b)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthor(a string) []book {
	return l[a]
}

func main() {
	// create a new library
	lib := make(library, 0)

	// add 2 books to the library by the same author
	book1 := book{
		title: "The Stand",
		author: author{
			name: "Stephen King",
		},
	}

	lib.addBook(book1)
	lib.addBook(book{title: "The Dark Tower", author: author{name: "Stephen King"}})

	// add 1 book to the library by a different author
	lib.addBook(book{title: "The Grapes of Wrath", author: author{name: "John Steinbeck"}})

	// dump the library with spew
	//spew.Dump(lib)

	// lookup books by known author in the library
	lbooks := lib.lookupByAuthor("Stephen King")
	//spew.Dump(lbooks)
	// print out the first book's title and its author's name
	//fmt.Print(lbooks[0].title, ", ", lbooks[0].author.name,"\n")
	if len(lbooks) != 0 {
		b := lbooks[0]
		fmt.Println(b.title, "by", b.author.name)
	}
}
