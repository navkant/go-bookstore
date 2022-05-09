package routes

import (
	"github.com/gorilla/mux"
	"github.com/navkant/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GE   T")
	router.HandleFunc("/book/{book_id}/", controllers.GETBookById).Methods("GET")
	router.HandleFunc("/book/{book_id}/", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{book_id}/", controllers.DeleteBook).Methods("DELETE")
}
