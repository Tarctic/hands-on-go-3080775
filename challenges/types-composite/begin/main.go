// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title string
	author
}

// define a library type to track a list of books
type library struct {
	authorBooks map[author][]book
}

// define addBook to add a book to the library
 func (l *library) addBook(a author, b []book) {
	l.authorBooks[a] = b
 }

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthor(authorName string) []book {
	return l.authorBooks[author{name: authorName}]
}

func main() {
	// create a new library
	l := library{
		authorBooks : map[author][]book{},
	}

	// add 2 books to the library by the same author
	author1 := author{
		name: "Veronica Roth",
	}

	book1 := book{
		title: "Divergent",
		author: author1,
	}
	
	book2 := book{
		title: "Carve the Mark",
		author: author1,
	}

	books1 := make([]book, 0)
	books1 = append(books1, book1, book2)
	l.addBook(author1, books1)
	
	// add 1 book to the library by a different author
	books2 := make([]book, 0)
	books2 = append(books2, book{title: "Rich Dad Poor Dad", author: author{name: "Robert Kiyosaki"}})
	l.addBook(author{name: "Robert Kiyosaki"}, books2)

	// dump the library with spew
	spew.Dump(l)

	// lookup books by known author in the library
	books := l.lookupByAuthor("Veronica Roth")

	// print out the first book's title and its author's name
	if len(books) != 0 {
		fmt.Println(books[0].title, books[0].author.name)
	}
}
