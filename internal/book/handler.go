package book

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Handler interface {
	GetBook(w http.ResponseWriter, r *http.Request)
	CreateBook(w http.ResponseWriter, r *http.Request)
}

type HandleBook struct {
	useCase UseCase
}

func NewHandlerBook(uc UseCase) Handler {
	return &HandleBook{useCase: uc}
}

func (h HandleBook) GetBook(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	resultChan := make(chan *Book)
	errChan := make(chan error)

	go func() {
		defer wg.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		book, err := h.useCase.Get(ctx, id)

		if err != nil {
			errChan <- err
			return
		}

		resultChan <- book

	}()

	select {
	case book := <-resultChan:
		if book != nil {
			json.NewEncoder(w).Encode(book)
			return
		}
		http.Error(w, "Book not fount", http.StatusNotFound)
	case err := <-errChan:
		log.Println("Error fetching book", err)
		http.Error(w, "Error fetching book", http.StatusInternalServerError)
	}
}

func (h HandleBook) CreateBook(w http.ResponseWriter, r *http.Request) {
	var bookRequest struct {
		Title       string `json:"title"`
		Genre       string `json:"genre"`
		Author      string `json:"author"`
		ReleaseDate string `json:"release_date"`
	}

	if err := json.NewDecoder(r.Body).Decode(&bookRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	releaseDate, err := time.Parse("2006-01-02", bookRequest.ReleaseDate)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultChan := make(chan int)
	errChan := make(chan error)

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		id, err := h.useCase.Create(ctx, bookRequest.Title, bookRequest.Genre, bookRequest.Author, releaseDate)

		if err != nil {
			errChan <- err
			return
		}
		resultChan <- id

	}()

	select {
	case id := <-resultChan:
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(strconv.Itoa(id)))
	case err := <-errChan:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
