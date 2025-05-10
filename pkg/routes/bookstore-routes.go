package routes

import (
	"net/http"

	"github.com/luksdan/go-bookstore/pkg/controllers"
)

func RegisterBookStoreRoutes() {
	http.HandleFunc("/book/", handleBooks)
	http.HandleFunc("/book/{id}", handleBook)
}

func handleBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		controllers.CreateBook(w, r)
	case "GET":
		controllers.GetBook(w, r)
	default:
		http.Error(w, "Method not alowed", http.StatusMethodNotAllowed)
	}
}

func handleBook(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		controllers.GetBookById(w, r)
	case "PUT":
		controllers.UpdateBook(w, r)
	case "DELETE":
		controllers.DeleteBook(w, r)
	default:
		http.Error(w, "Method not alowrd", http.StatusMethodNotAllowed)
	}
}
