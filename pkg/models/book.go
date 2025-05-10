package models

import (
	"database/sql"
	"log"

	"github.com/luksdan/go-bookstore/pkg/config"
)

var db *sql.DB

type Book struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func (b *Book) CreateBook() *Book {
	query := `INSERT INTO books(
			  name, author, publication) 
			  VALUES (
			  ?,?,?)`
	_, err := db.Exec(query, b.Name, b.Author, b.Publication)
	if err != nil {
		log.Fatal("Error inserting table")
	}
	return b
}

func GetAllBooks() []Book {
	query := `SELECT id, name, author, publication
		 	  FROM books`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error selecting", err)
	}
	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.Id, &book.Name, &book.Author, &book.Publication)
		if err != nil {
			log.Fatal("Error scaning row", err)
		}
		books = append(books, book)
	}
	return books
}

func GetBookById(Id int64) *Book {
	db = config.GetDB()
	query := `SELECT id, name, author, publication
			  FROM books
			  WHERE id = ?`
	row := db.QueryRow(query, Id)
	var book Book
	err := row.Scan(&book.Id, &book.Name, &book.Author, &book.Publication)
	if err != nil {
		log.Fatal("Error selecting by id")
	}
	return &book
}

func DeleteBook(Id int64) Book {
	var book Book
	querySelect := `SELECT name, author, publication
					FROM books
					WHERE id = ?`
	row := db.QueryRow(querySelect, Id)
	err := row.Scan(&book.Name, &book.Author, &book.Publication)
	if err != nil {
		log.Fatal("Error trying to scan", err)
	}

	queryDelete := `DELETE FROM books
			  WHERE id = ?`
	_, errDelete := db.Exec(queryDelete, Id)
	if errDelete != nil {
		log.Fatal("Error deleting record")
	}
	return book
}

func (book *Book) UpdatBook(Id int64) Book {
	query := `UPDATE books SET
			  name = ?, author = ?, publication = ?
			  WHERE id = ?`
	_, err := db.Query(query, book.Name, book.Author, book.Publication, Id)
	if err != nil {
		log.Fatal("Error trying to update")
	}
	return *book
}
