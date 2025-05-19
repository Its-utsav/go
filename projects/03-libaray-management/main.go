package main

import (
	"fmt"
	"time"
)

type Book struct {
	Id              int
	Title           string
	Author          *Author
	Publisher       string
	PublicationDate time.Time
}

func newBook(id int, title string, author *Author, publisher string) *Book {
	book := Book{
		Id:        id,
		Title:     title,
		Author:    author,
		Publisher: publisher,
	}
	return &book
}

func (book *Book) getBookInfo() {
	fmt.Printf("Title : %s , Author : %s, Publisher : %s ,Publication Date : %s", book.Title, book.Author.getAuthorName(), book.Publisher, book.PublicationDate)
}
func (book *Book) updatePublisher(newPublisher string) {
	book.Publisher = newPublisher
}

type Author struct {
	Id   int
	Name string
	Bio  string
}

func newAuthor(id int, name, bio string) *Author {
	author := Author{
		Id:   id,
		Name: name,
		Bio:  bio,
	}
	return &author
}

func (author *Author) getAuthorName() string {
	return author.Name
}
func (author *Author) updateAuthorBio(newBio string) {
	author.Bio = newBio
}

func main() {
	me := newAuthor(1, "Utsav", "I am Learning Go Lang")
	jsForChild := newBook(101, "jsForChild", me, "me")
	// fmt.Println(jsForChild)
	jsForChild.getBookInfo()
}
