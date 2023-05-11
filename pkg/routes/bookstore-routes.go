package routes

import (
	"github.com/Sanjay-21pw/go-bookstore-new/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.createBook).Methods("POST")
	router.HandleFunc("/book/", controllers.getBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.getBookById).Methods("GET")
	router.HandleFunc("book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.deleteBook).Methods("DELETE")

}
