package main

import (
	"github/phses/go-crud-api/internal/book"
	"github/phses/go-crud-api/internal/repository/db"
	"github/phses/go-crud-api/internal/repository/postgres"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	db := db.NewPostgresDb()

	bookRepo := postgres.NewPostgresRepository(db)
	bookUseCase := book.NewBookUseCase(bookRepo)
	bookHandler := book.NewHandlerBook(bookUseCase)

	r := mux.NewRouter()
	r.HandleFunc("/books", bookHandler.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", bookHandler.GetBook).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}
