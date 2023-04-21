package main

import (
	"fmt"
	"time"
)

type Title string
type Name string

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int
	lended int
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Printf(string(member.name), title, audit.checkOut.String(), "through", returnTime)
	}
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMeberAudits(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)

	}
}

func printLibraryBooks(library *Library) {
	for title, entry := range library.books {
		fmt.Printf(string(title), entry.total, entry.lended)
	}
}

func checkoutBook(library *Library, member *Member, title Title) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not found")
		return false
	}

	if book.lended == book.total {
		fmt.Println("Book not available")
		return false
	}
	book.lended++
	library.books[title] = book

	member.books[title] = LendAudit{checkOut: time.Now()}
	return true
}

func returnBook(library *Library, member *Member, title Title) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not found")
		return false
	}

	audit, found := member.books[title]
	if !found {
		fmt.Println("Book not checked out")
		return false
	}

	book.lended--
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}

func main() {
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	library.books["Webapps"] = BookEntry{total: 10, lended: 0}
	library.books["Go"] = BookEntry{total: 5, lended: 0}

	library.members["Miha"] = Member{name: "Miha", books: make(map[Title]LendAudit)}

	fmt.Println("Books in library:")
	printLibraryBooks(&library)
	printMeberAudits(&library)

	member := library.members["Miha"]
	checkedOut := checkoutBook(&library, &member, "Webapps")
	fmt.Println("Books in library:")
	if checkedOut {
		printLibraryBooks(&library)
		printMeberAudits(&library)
	}
}
